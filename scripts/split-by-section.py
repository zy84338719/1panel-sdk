#!/usr/bin/env python3
"""Split a service file at `// === Section Name ===` boundaries."""
import re
import sys
from pathlib import Path


def split_file(src: Path, mapping: dict[str, str]):
    """Split `src` into multiple files based on `// === Section ===` markers.

    `mapping` is a dict from section name (the text between `===` and `===`)
    to the output file name. Methods not in any section go to `default`.
    The type declaration, package/imports header and `Call()` method are
    always included at the top of every output file.
    """
    text = src.read_text()
    # Find package + imports header.
    m = re.search(r'^(//.*\n)*package onepanel\n+(import [^\n]+\n+(\nimport \([^)]+\)\n+)?)?',
                  text, re.MULTILINE)
    if not m:
        # Try the simpler pattern (just package + import).
        m = re.search(r'^package onepanel\n', text, re.MULTILINE)
    header_end = m.end() if m else 0
    header = text[:header_end]
    body = text[header_end:]

    # Split body at section markers. Sections look like:
    #   // === Container CRUD ===
    section_re = re.compile(r'^// === (.+?) ===\n', re.MULTILINE)

    # Walk through body, group methods by section.
    parts = re.split(section_re, body)
    # parts[0] is preamble (no section), parts[1] is first section name,
    # parts[2] is first section body, parts[3] is next section name, etc.
    sections = {}
    preamble = parts[0] if parts else ''
    i = 1
    while i < len(parts):
        name = parts[i].strip()
        content = parts[i + 1] if i + 1 < len(parts) else ''
        sections[name] = content
        i += 2

    # If 'preamble' contains method bodies (database.go's MySQL section
    # has no `// === MySQL ===` header) and the mapping defines a
    # catch-all target, treat preamble as a section under that name.
    if preamble.strip():
        for section_name, fname in mapping.items():
            if fname.endswith('mysql.go') and 'MySQL' not in sections:
                # Strip the trailing blank line and feed preamble in as
                # the MySQL section content.
                sections['MySQL'] = preamble.rstrip() + '\n\n'
                preamble = ''
                break

    # Header: type + Call() are always at the top of body, BEFORE the
    # first `// ===` marker. Find them in preamble.
    type_call = preamble.strip() + '\n\n'

    # Group sections into output files. The type def + Call() go in
    # only the first output file (alphabetical) — Go does not allow
    # `type X struct {}` to be repeated across files, but method
    # declarations on a receiver ARE allowed to live in any file of
    # the same package.
    sorted_files = sorted(set(mapping.values()))
    outputs: dict[str, list[str]] = {fname: [header] for fname in sorted_files}
    outputs[sorted_files[0]].append(type_call)  # type + Call only here.
    for section, content in sections.items():
        fname = mapping.get(section, mapping.get('__default__'))
        if fname is None:
            continue
        outputs[fname].append(f'// === {section} ===\n{content}')

    # Ensure each output file has type + Call.
    written = []
    for fname, lines in outputs.items():
        path = Path(fname)
        # Always overwrite — the original src will be deleted below.
        path.write_text(''.join(lines))
        written.append(fname)
    return written


if __name__ == '__main__':
    target = sys.argv[1] if len(sys.argv) > 1 else None
    if target == 'container':
        mapping = {
            'Container CRUD': 'container_crud.go',
            'Container files': 'container_files.go',
            'Image registry / repo': 'container_repo.go',
            'Compose': 'container_compose.go',
            'Compose templates': 'container_compose_template.go',
            'Image': 'container_image.go',
            'Network': 'container_network.go',
            'Volume': 'container_volume.go',
            'Daemon JSON / docker status': 'container_daemon.go',
        }
        src = Path('container.go')
    elif target == 'database':
        # MySQL methods have no `// ===` section header — they live
        # between the type def and the first `// === Redis ===` marker.
        # Detect them as the "default" / "MySQL" pre-section.
        mapping = {
            'MySQL': 'database_mysql.go',                # catch-all default
            'Redis': 'database_redis.go',
            'DB generic (per-database engine)': 'database_common.go',
            'PostgreSQL': 'database_postgres.go',
            'MongoDB': 'database_mongodb.go',
        }
        src = Path('database.go')
    else:
        print('usage: split-by-section.py <container|database>')
        sys.exit(1)
    written = split_file(src, mapping)
    src.unlink()
    for f in written:
        print(f'wrote {f}')
