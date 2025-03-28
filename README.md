# GoMapper - Port Scanner in GO 

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.24-blue?logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/github/license/3rr0r-505/GoMapper)](https://github.com/3rr0r-505/GoMapper/blob/main/LICENSE)
[![Release Version](https://img.shields.io/github/v/release/3rr0r-505/GoMapper?label=latest)](https://github.com/3rr0r-505/GoMapper/releases)
[![Issues](https://img.shields.io/github/issues/3rr0r-505/GoMapper)](https://github.com/3rr0r-505/GoMapper/issues)

<pre>
    ________     ______  ___                                   
  __  ____/________   |/  /_____ _____________________________
  _  / __ _  __ \_  /|_/ /_  __ `/__  __ \__  __ \  _ \_  ___/
  / /_/ / / /_/ /  /  / / / /_/ /__  /_/ /_  /_/ /  __/  /    
  \____/  \____//_/  /_/  \__,_/ _  .___/_  .___/\___//_/     
                                 /_/     /_/             _ 
</pre>

**GoMapper** is a fast and efficient tool for scanning ports on a target system, similar to Nmap. It allows you to scan specific ports or a range of ports to gather information about open services.

## Features

- Scan specific ports or port ranges.
- Support for different platforms (Windows, Linux).
- Banner grabbing for open ports.
- Fast and easy to use.
- Written in Go (No need for Go installation on user machines).

## Installation

### Install with Go

If you have Go installed, you can easily install **GoMapper** using the following command:

```bash
go install github.com/3rr0r-505/GoMapper@latest
```

This will download and install GoMapper into your `$GOPATH/bin` (or `$GOBIN`) directory. Once installed, you can run GoMapper from anywhere in your terminal.

Alternatively, if you want to install a specific version, you can do so like this:

```bash
go install github.com/3rr0r-505/GoMapper@<release tag>
```

### Update with Go
```bash
rm -rf $(go env GOPATH)/pkg/mod/github.com/3rr0r-505
rm -f $(go env GOPATH)/bin/GoMapper
go clean -cache && go clean -modcache
go install github.com/3rr0r-505/GoMapper@<release tag>
```

### Install from Release

You can also download pre-built executables for different platforms directly from the [releases page](https://github.com/3rr0r-505/GoMapper/releases).

Windows 64-bit: `GoMapper-windows-amd64.exe`

Windows 32-bit: `GoMapper-windows-x86.exe`

Linux 64-bit: `GoMapper-linux-amd64`

Linux 32-bit: `GoMapper-linux-x86`

Download the appropriate binary for your platform, make it executable, and run it directly.

## Usage
After installation, you can run GoMapper from your terminal using the following command:
```bash
GoMapper --help
```

### Command Options

`--host`: Specify the target host or IP address (default: 127.0.0.1).

`--ports`: Specify a comma-separated list of ports or a range of ports (default: the top 1000 ports from Nmap).

`--output`: Optionally, store the result in a text file.

### Examples
Scan a specific host and ports:
```bash
GoMapper --host 192.168.1.1 --ports 80,443
```

Scan a range of ports on a host:
```bash
GoMapper --host example.com --ports 20-25
```

Scan all ports and save the output to a file:
```bash
GoMapper --host example.com --ports * --output "scan.txt"
```

---

## Contributing
Contributions are welcome! Feel free to open issues or pull requests for any improvements or bug fixes. For more info, visit [CONTRIBUTING.md](https://github.com/3rr0r-505/GoMapper/blob/main/CONTRIBUTING.md)

## Legal Disclaimer
The use of code contained in this repository, either in part or in its totality,
for engaging targets without prior mutual consent is illegal. **It is
the end user's responsibility to obey all applicable local, state and
federal laws.**

Developers assume **no liability** and are not
responsible for misuses or damages caused by any code contained
in this repository in any event that, accidentally or otherwise, it comes to
be utilized by a threat agent or unauthorized entity as a means to compromise
the security, privacy, confidentiality, integrity, and/or availability of
systems and their associated resources. In this context the term "compromise" is
henceforth understood as the leverage of exploitation of known or unknown vulnerabilities
present in said systems, including, but not limited to, the implementation of
security controls, human- or electronically-enabled.

The use of this code is **only** endorsed by the developers in those
circumstances directly related to **educational environments** or
**authorized penetration testing engagements** whose declared purpose is that
of finding and mitigating vulnerabilities in systems, limiting their exposure
to compromises and exploits employed by malicious agents as defined in their
respective threat models.

## License
This project is licensed under the [MIT License](https://github.com/3rr0r-505/GoMapper/blob/main/LICENSE) - see the LICENSE file for details.
