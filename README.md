# protobuf-test-202401

```diff
diff v1.A -> v2.B
  any(Inverse(protocmp.Transform, protocmp.Message{
-       "@type": s"v1.A",
+       "@type": s"v2.B",
        "a":     string("a"),
        "b": protocmp.Message{
-               "@type": s"v1.SuperSet",
+               "@type": s"v2.SuperSet",
                "first": int32(100),
                "third": int32(300),
        },
        "c": []protocmp.Message{
                {
-                       "@type": s"v1.SuperSet",
+                       "@type": s"v2.SuperSet",
                        "first": int32(100),
                        "third": int32(300),
                },
                {
-                       "@type": s"v1.SuperSet",
+                       "@type": s"v2.SuperSet",
                        "first": int32(101),
                        "third": int32(301),
                },
        },
        "d": []protocmp.Message{
                {
-                       "@type": s"v1.Wrapper",
+                       "@type": s"v2.Wrapper",
                        "set": protocmp.Message{
-                               "@type": s"v1.SuperSet",
+                               "@type": s"v2.SuperSet",
                                "first": int32(102),
                                "third": int32(302),
                        },
                },
        },
  }))

diff v2.B -> v1.A
  any(Inverse(protocmp.Transform, protocmp.Message{
-       "@type": s"v2.B",
+       "@type": s"v1.A",
        "a":     string("a"),
        "b": protocmp.Message{
+               "2":      protoreflect.RawFields{0x10, 0xc8, 0x01},
-               "@type":  s"v2.SuperSet",
+               "@type":  s"v1.SuperSet",
                "first":  int32(100),
-               "second": int32(200),
                "third":  int32(300),
        },
        "c": []protocmp.Message{
                {
+                       "2":      protoreflect.RawFields{0x10, 0xc8, 0x01},
-                       "@type":  s"v2.SuperSet",
+                       "@type":  s"v1.SuperSet",
                        "first":  int32(100),
-                       "second": int32(200),
                        "third":  int32(300),
                },
                {
+                       "2":      protoreflect.RawFields{0x10, 0xc9, 0x01},
-                       "@type":  s"v2.SuperSet",
+                       "@type":  s"v1.SuperSet",
                        "first":  int32(101),
-                       "second": int32(201),
                        "third":  int32(301),
                },
        },
        "d": []protocmp.Message{
-               s"{set:{first:102  third:302}  fourth:402}",
+               s"{set:{first:102  third:302}  2:402}",
        },
  }))

a v1.A
{
  "a":  "a",
  "b":  {
    "first":  100,
    "third":  300
  },
  "c":  [
    {
      "first":  100,
      "third":  300
    },
    {
      "first":  101,
      "third":  301
    }
  ],
  "d":  [
    {
      "set":  {
        "first":  102,
        "third":  302
      }
    }
  ]
}

b v1.B
{
  "a":  "a",
  "b":  {
    "first":  100,
    "second":  200,
    "third":  300
  },
  "c":  [
    {
      "first":  100,
      "second":  200,
      "third":  300
    },
    {
      "first":  101,
      "second":  201,
      "third":  301
    }
  ],
  "d":  [
    {
      "set":  {
        "first":  102,
        "third":  302
      },
      "fourth":  402
    }
  ]
}

a2 v1.A
{
  "a":  "a",
  "b":  {
    "first":  100,
    "third":  300
  },
  "c":  [
    {
      "first":  100,
      "third":  300
    },
    {
      "first":  101,
      "third":  301
    }
  ],
  "d":  [
    {
      "set":  {
        "first":  102,
        "third":  302
      }
    }
  ]
}
```
