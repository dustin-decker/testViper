Output from Viper:
```
(main.Config) {
 ListenInterface: (string) (len=9) "127.0.0.1",
 ListenPort: (int) 8080,
 Thing: (map[string]main.SomeOtherThing) (len=3) {
  (string) (len=11) "fully-works": (main.SomeOtherThing) {
   Stuff: (string) (len=5) "works"
  },
  (string) "": (main.SomeOtherThing) {
   Stuff: (string) ""
  },
  (string) (len=9) "partially": (main.SomeOtherThing) {
   Stuff: (string) ""
  }
 }
}
```

Output from go-yaml:
```
(main.Config) {
 ListenInterface: (string) (len=9) "127.0.0.1",
 ListenPort: (int) 8080,
 Thing: (map[string]main.SomeOtherThing) (len=3) {
  (string) (len=15) "partially.works": (main.SomeOtherThing) {
   Stuff: (string) (len=7) "missing"
  },
  (string) (len=11) "fully-works": (main.SomeOtherThing) {
   Stuff: (string) (len=5) "works"
  },
  (string) (len=7) ".broken": (main.SomeOtherThing) {
   Stuff: (string) (len=7) "missing"
  }
 }
}
```
