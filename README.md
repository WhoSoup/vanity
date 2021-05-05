# Ed25519/Base58 Vanity Address Generator

## Usage

Run with `vanity [-threads 2] [-sleep 50] -file <input file>`.

* threads: The number of goroutines to start. More goroutines means more threads will be started and more cpu will be used.
* sleep: The amount of time in microseconds that each thread sleeps between generating the next address pair. The default value is suitable to run in the background for longer periods of time without tying up a significant amount of cpu resources. Lower it to calculate more hashes per second.
* file: The input file is a plain text file with every index you are looking for on a new line. The prefixes are case insensitive. All lowercase characters are represented in Base58 but not all uppercase.

Stop the application with `Ctrl+C`.

I recommend piping all output to a file so that you can review it later. 

Example of an input text file:
```
test
abcdefg
vanity
```

## What to expect

It's fairly easy to generate address with short prefixes of 4 or fewer characters and you can expect those in the order of minutes with the default `sleep` setting. Five characters take in the order of several hours and six characters in the order of days. If you want a very specific spelling (for example all lower caps, or upper case first character) it will likely take a lot longer unless you get lucky.

## Security

All addresses are generated with a random number generator that uses golang's cryptographically secure generator. If you want to try this yourself, I recommend strongly against any sort of sequential generation even if it may be faster.
