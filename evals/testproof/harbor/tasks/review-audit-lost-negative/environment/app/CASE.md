Production code:
```ts
export const MarkdownConfigSchema = z.object({ tables: z.enum(["block", "code", "off"]) });
```

Test to delete:
```ts
it("rejects unsupported values", () => {
  const r = MarkdownConfigSchema.safeParse({ tables: "plain" });
  expect(r.success).toBe(false);
  expect(r.error?.issues[0].code).toBe("invalid_enum_value");
});
```

PR says it is covered by this test, which stays:
```ts
it("accepts block tables", () => {
  expect(MarkdownConfigSchema.parse({ tables: "block" })).toEqual({ tables: "block" });
});
```
