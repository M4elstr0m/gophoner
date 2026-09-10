# gophoner​​     <img src="https://github.com/user-attachments/assets/82a84627-3382-405c-8972-05f2b328559d" alt="Gophoner icon" width="64" height="64" align="middle">

Check simultaneously if a phone number is registered on popular apps & websites, without any prerequisite.

![Visitor Badge](https://visitor-badge.laobi.icu/badge?page_id=M4elstr0m.gophoner&left_text=Visitors&right_color=orange)
![Stars Badge](https://img.shields.io/github/stars/M4elstr0m/gophoner?style=flat&color=yellow&label=Stars)
![License Badge](https://img.shields.io/badge/License-PolyForm%20Internal%20Use%20License%201.0.0-lightgrey)
![Go Version Badge](https://img.shields.io/badge/Go-26.0%2B-blue)

Please star & share this repository if you find it useful :star:

![Gophoner animation](https://github.com/user-attachments/assets/1fb7b01b-3953-40ba-88ad-28fd416a37a1)

## Table of Contents

- [Disclaimer](#disclaimer-️warning)
- [Features](#features)
- [Installation](#installation-computer)
- [Usage](#usage-books)
- [Supported Modules](#modules-jigsaw)
- [Roadmap](#roadmap-world_map)
- [Legal & Ethical Use](#legal--ethical-use-scroll)
- [License](#license-balance_scale)
- [Credits](#credits-newspaper)

## Disclaimer :warning:

> [!CAUTION]
> It is important you read the full policy of this project before using or contributing to this tool.
> 
> Please refer to the following sections:
> - [Legal & Ethical Use](#legal--ethical-use-scroll)
> - [License](#license-balance_scale)

## Features

![Gophoner in-action](https://github.com/user-attachments/assets/6c0d27e1-2bab-4d23-9063-3ecfb9882657)

- **CLI and TUI**: use `gophoner check` for scripting and automation, or `gophoner interactive` for a guided terminal interface.
- **Cross-platform**: prebuilt binaries for Windows, Linux, and macOS.
- **Concurrent module checks**: every selected module runs in parallel, so checking a phone number against several services takes about as long as the slowest one, not the sum of them all.
- **Silent by default**: none of the currently supported modules alert or notify the target phone number, no SMS or email triggered while checking. If a future module does trigger a notification on the target's end, it will be called out explicitly.
- **Privacy conscious logging**: target phone numbers are never written to the log file unless `--debug` is set, and `--no-log` disables all logging entirely.
- **Randomized browser fingerprinting**: each request gets an internally consistent TLS, User-Agent, and Client Hints profile drawn from a real desktop browser pool, making requests far harder to fingerprint and block than a plain HTTP client.
- **Single target scope**: checks one phone number against your chosen modules per invocation, not built for bulk enumeration.
- **Update aware**: checks for new releases on startup and lets you know when one is available, disable with `--no-update`.

## Installation :computer:

![go install](https://img.shields.io/badge/Go%20Install-00ADD8?logo=Go&logoColor=white&style=for-the-badge)

Using `go install` (Requires [Go](https://go.dev/dl/) 1.26 or newer)
```sh
go install github.com/M4elstr0m/gophoner/cmd/gophoner@latest
```

![arch aur](https://img.shields.io/badge/AUR%20(PARU%20/%20YAY)-1793D1?logo=ArchLinux&logoColor=white&style=for-the-badge)

> [!IMPORTANT]
> This only works on Arch-based Linux distributions.

```sh
# WORK IN PROGRESS (THIS IS PLANNED)
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

## Modules :jigsaw:

**v1.0.0 Modules**

![Amazon](https://custom-icon-badges.demolab.com/badge/Amazon-FF6201?logo=Amazon&logoColor=white&style=for-the-badge)
![Microsoft](https://custom-icon-badges.demolab.com/badge/Microsoft-89d2ff?logo=Microsoft&logoColor=black&style=for-the-badge)
![Facebook](https://img.shields.io/badge/Facebook-0866FF?logo=Facebook&logoColor=white&style=for-the-badge)
![Google](https://img.shields.io/badge/Google-4285F4?logo=Google&logoColor=white&style=for-the-badge)
![OpenAI](https://custom-icon-badges.demolab.com/badge/OpenAI-74aa9c?logo=OpenAI&logoColor=white&style=for-the-badge)
<!-- 
![Telegram]
![Signal]
![Instagram](https://img.shields.io/badge/Instagram-FF0069?logo=Instagram&logoColor=white&style=for-the-badge)
![Snapchat](https://img.shields.io/badge/Snapchat-FFFC00?logo=Snapchat&logoColor=black&style=for-the-badge)
-->

## Roadmap :world_map:

- [ ] v1.0.0
  - [x] Initial tool and Amazon module
  - [x] More modules: Microsoft, Facebook, Amazon & plenty more surprises :no_mouth:
  - [x] CLI QoL: progress bar
  - [x] **gophoner** logo & a see it in action GIF
- [ ] CLI QoL: only display positive results, progressive results display, fun facts during progress
- [ ] Redirect output to JSON
- [ ] More modules!
- [ ] A plug-in system (maybe)

## Legal & Ethical Use :scroll:

**gophoner** is built for security research, personal OSINT hygiene (checking
your own digital footprint), and educational purposes.

Checking a phone number against a service without the number holder's
consent may violate that service's Terms of Service, and, depending on
your jurisdiction and how you use the results, local privacy, stalking,
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

## Credits :newspaper:

- The guy who tried to log into one of my accounts at 3 AM, 20 times in 30 minutes, then gave up. He gave me the insomnia that built this tool.
- [sherlock](https://github.com/sherlock-project/sherlock) & [holehe](https://github.com/megadose/holehe), tools that also use a single credential to check hundreds of services
- [ignorant](https://github.com/megadose/ignorant), the first tool I found when searching for something that already did what I had in mind

---

[Go back to top :arrow_up:](#gophoner)
