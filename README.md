# What's Checked Out?

The public web interface for Biblionix is, frankly, trash. And it's even worse when you have multiple accounts.  Checking what we have checked out with a family of five is a chore.

Luckily, it operates over a loosely secured API, so we can automate it!

That's what "What's Checked Out?" does.  It takes a Biblionix account, password and subdomain, logs in, gets the list of checked out items, and repeats for every linked account as well.

At the end you get a summary of everything that is currently out.

## Example

```bash
$ cat config
account 12345
password 8675309
subdomain blair

$ ./whats-checked-out -config config

=== John (Adult) [12345] ==============

Real Japanese cooking: the only Japanese cookbook you will ever need! : traditions, tips & techniques with over 600 authentic recipes | Itoh, Makiko | 11‑11‑2025

=== Darcy (Adult) [12346] ==============

Unleashed: Dog Man — bk 2                                            | Pilkey, Dav (1966-) | 11‑4‑2025
The big book of amazing LEGO creations: with bricks you already have | Dees, Sarah,        | 11‑11‑2025

$
```
