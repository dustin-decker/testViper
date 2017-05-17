Output looks like this:
```
(main.Config) {
 ListenInterface: (string) (len=9) "127.0.0.1",
 ListenPort: (int) 8080,
 Thing: (map[string]main.SomeOtherThing) (len=3) {
  (string) "": (main.SomeOtherThing) {
   Stuff: (string) ""
  },
  (string) (len=9) "partially": (main.SomeOtherThing) {
   Stuff: (string) ""
  },
  (string) (len=11) "fully-works": (main.SomeOtherThing) {
   Stuff: (string) (len=5) "works"
  }
 }
}
```
