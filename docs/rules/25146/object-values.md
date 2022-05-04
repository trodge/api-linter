--------------------------------------------------------------------------------

rule: aip: 25146 name: [core, '25146', object-values] summary: Avoid maps with
object values. permalink: /25146/standardized-codes redirect_from:

## - /25146/standardized-codes

# Any

This rule discourages the use of maps with object values, as described in
[AIP-25146][].

## Details

This rule complains if it sees a map field with objects as values. Output-only
fields are excluded.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message Book {
  // Maps with object values are discouraged.
  map<int32, Page> contents = 1;
}
```

**Correct** code for this rule:

The correct code is likely to vary substantially by use case. See [AIP-25146][]
for details and tradeoffs of various approaches for generic fields.

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aip.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::25146::object-values=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)
message Book {
  map<int32, Page> contents = 1;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aip-25146]: https://aip.dev/25146
[aip.dev/not-precedent]: https://aip.dev/not-precedent
