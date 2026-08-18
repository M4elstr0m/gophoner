# gophoner

Check simultaneously if a phone number is registered on popular apps & websites

![Visitor Badge](https://visitor-badge.laobi.icu/badge?page_id=M4elstr0m.gophoner&left_text=Visitors&right_color=orange)
![Stars Badge](https://img.shields.io/github/stars/M4elstr0m/gophoner?style=flat&color=yellow&label=Stars)
![License Badge](https://img.shields.io/badge/License-PolyForm%20Internal%20Use%20License%201.0.0-lightgrey)
![Go Version Badge](https://img.shields.io/badge/Go-26.0%2B-blue)

Please star & share this repository if you find it useful ⭐

## Disclaimer

> [!CAUTION]
> It is important you read the full policy of this project before using or contributing to this tool.
> 
> Please refer to the following sections:
> * [Legal & Ethical Use](#legal--ethical-use-)
> * [License](#license-️)

## Installation :computer:

Using `go install` (Requires [Go](https://go.dev/dl/) 1.26 or later.)
```sh
go install github.com/M4elstr0m/gophoner/cmd/gophoner@latest
```

---

Check if the installation succeeded
```sh
gophoner version
```

## Usage :books:

```sh
gophoner help

gophoner interactive

gophoner check -t +15551234567 -A
```

## Legal & Ethical Use :scroll:

**gophoner** is built for security research, personal OSINT hygiene (checking
your own digital footprint), and educational purposes.

Checking a phone number against a service without the number holder's
consent may violate that service's Terms of Service, and — depending on
your jurisdiction and how you use the results — local privacy, stalking,
or harassment laws. You are solely responsible for ensuring your use of
this tool is lawful and authorized.

Do not use **gophoner** to harass, stalk, dox, or otherwise target an
individual without their consent. The author provides this software as
is, disclaims all liability for its use or misuse, and does not endorse
or support any unlawful use of it.

**gophoner** is designed with the operator's own privacy in mind: by
default it never writes target phone numbers to its log file (only at
`--debug` level), and a `--no-log` is also at your disposal to disable all logging entirely.

**gophoner** is also architecturally scoped to reduce abuse potential: it
checks a single phone number against your chosen modules per invocation.
It is not built for bulk enumeration or sweeping ranges of numbers, even though a basic script wrapping the CLI could still automate that.

## License :balance_scale:

Licensed under **PolyForm Internal Use License 1.0.0**.

Here is a plain-language summary, which is not a substitute for the license itself:

**:green_square: You can:**
- Use **gophoner**, including for your company's internal business operations, commercial use included.
- Modify the code for your own private/internal use.
- Fork it to prepare and submit a pull request back to this repository.
- Read the full source code.

**:red_square: You can't:**
- Distribute or redistribute **gophoner**, or any modified version of it, to anyone else.
- Offer it as a hosted or managed service to third parties.
- Sublicense or transfer your rights to someone else.
- Expect a warranty: it's provided "as is", with no liability to the author.

If **gophoner** is useful to you, a star, or a mention to your colleagues, is always appreciated :wink: