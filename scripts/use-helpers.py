#!/usr/bin/env python3
"""Rewrite hand-written typed methods to use ServiceBase helper methods.

Before:
    var out map[string]any
    if err := s.Get(ctx, "/containers/status", &out); err != nil {
        return nil, err
    }
    return out, nil

After:
    return s.getMap(ctx, "/containers/status")

Skips:
  - services_swagger.go (generated)
  - *_test.go files
  - any method whose body uses `out` after the helper call
"""
import re
import sys
from pathlib import Path

HELPER = {'Get': 'getMap', 'Post': 'postMap', 'Put': 'putMap', 'Delete': 'deleteMap'}
SKIP = {'services_swagger.go', 'doc.go'}

# Pattern: matches the 4-line call pattern with optional body.
# Captures: (verb, path_lit, body_or_empty)
PATTERN = re.compile(
    r'\tvar out map\[string\]any\n'
    r'\tif err := s\.(Get|Post|Put|Delete)\(ctx, ([^,]+(?:, [^,]+)*?)(?:, (.*?))?, &out\); err != nil \{\n'
    r'\t\treturn nil, err\n'
    r'\t\}\n'
    r'\treturn out, nil',
    re.MULTILINE,
)

# Path literal can be a string or fmt.Sprintf(...) — captured.

def rewrite(path: Path) -> int:
    text = path.read_text()
    count = 0

    def replace(m):
        nonlocal count
        verb = m.group(1)
        path_lit = m.group(2).strip()
        body_lit = m.group(3)
        helper = HELPER[verb]
        if body_lit is None:
            call = f'return s.{helper}(ctx, {path_lit})'
        else:
            body_lit = body_lit.strip()
            call = f'return s.{helper}(ctx, {path_lit}, {body_lit})'
        count += 1
        return call

    new_text = PATTERN.sub(replace, text)
    if count:
        path.write_text(new_text)
    return count


def main():
    files = [p for p in Path('.').glob('*.go')
             if p.name not in SKIP and not p.name.endswith('_test.go')]
    total = 0
    for p in files:
        n = rewrite(p)
        if n:
            print(f'{p}: {n} method(s) rewritten')
        total += n
    print(f'\ntotal: {total} method(s) rewritten across {len(files)} files')


if __name__ == '__main__':
    main()
