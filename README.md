diceware
========

This tool is simple enough that you can audit the code yourself. It relies on
golang's crypto/rand for entropy, which tries to get random bits from truly
random sources as much as possible. The tool is intended to generate passwords
locally on your own computer using the diceware wordlist from:

http://world.std.com/~reinhold/diceware.html

...specifically:

http://world.std.com/%7Ereinhold/diceware.wordlist.asc

I use the tool to generate pass phrases which I then store in LastPass (since
LastPass does not generate pass phrases on its own). For my master password, I
used the actual rolling of a physical die because this program could leave a
residue of the password you generated on your computer, or it might be
comprimised by a hacker (I mean, if you are _really_ paranoid) so it should not
be relied upon for passwords that grant broad access to all your services (e.g.,
your LastPass or KeePass master password). However, I think it is reasonable to
use this program to generate passwords for individual services which you then
store in a password manager.

To use:

```
go build
./diceware --help
./dicware
```

By default it generates a 7-word pass phrase but you can change the number of
words it uses. The included dictionary is from world.std.com but I removed the
PGP stuff at the beginning and end, otherwise the program cannot parse the
dictionary.

Enjoy.
