### Installing Dependencies

> Before installing Sponge, you need to install two dependencies: `go` and `protoc`.

**✅ Installing go**

Download go from: [https://go.dev/doc/install](https://go.dev/doc/install)

> Required version: 1.16 or above. Add the directory containing the binary files generated by the `go install` command (usually $GOPATH/bin) to the system environment variable "path".

Check go version: `go version`

<br>

**✅ Installing protoc**

Download protoc from: [https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.3)

> Required version: v3.20 or above. Add the directory containing the protoc binary file (recommended: $GOPATH/bin) to the system environment variable "path".

Check protoc version: `protoc --version`

<br>

After installing go and protoc, proceed to install Sponge. Sponge can be installed on Windows, macOS, and Linux environments.

<br>
<br>

### Linux or macOS Environment

```bash
# Install Sponge
go install github.com/zhufuyi/sponge/cmd/sponge@latest

# Initialize Sponge, automatically install Sponge's dependency plugins
sponge init

# Check if all plugins have been successfully installed. If any plugins fail to install, retry with the command: sponge tools --install
sponge tools

# Check Sponge version
sponge -v
```

<br>
<br>

### Windows Environment

> In the Windows environment, you need to install mingw64, make, and cmder to support the Linux command environment required by Sponge.

**✅ Installing mingw64**

Download mingw64 from: [x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z](https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/seh/x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z)

After downloading, extract it to the directory `D:\Program Files\mingw64`, and add the directory containing frequently used Linux commands, such as `D:\Program Files\mingw64\bin`, to the system environment variable "PATH".

<br>

**✅ Installing the make command**

Navigate to the directory `D:\Program Files\mingw64\bin`, find the executable file `mingw32-make.exe`, and copy and rename it to `make.exe`.

Check make version: `make -v`

<br>

**✅ Installing cmder**

Download cmder from: [cmder-v1.3.20.zip](https://github.com/cmderdev/cmder/releases/download/v1.3.20/cmder.zip)

After downloading, extract it to the directory `D:\Program Files\cmder`, and add the directory `D:\Program Files\cmder` to the system environment variable "PATH".

Perform basic configuration for cmder:

- **Resolve Space Issue While Typing Commands**: Open the cmder interface, press `win+alt+p` to access the settings, search for `Monospace` in the upper left corner, uncheck it, and save and exit.
- **Configure Right-Click Launch for cmder**: Press the `win+x` key combination, then press the letter `a` to enter the administrative terminal. Execute the command `Cmder.exe /REGISTER ALL`. Right-click in any folder and choose `Cmder Here` to open the cmder interface.

> ⚠ When developing projects with Sponge in a Windows environment, to avoid issues with missing Linux commands, please use cmder instead of the built-in cmd terminal, the terminal in Goland, or the terminal in VS Code.

Open the `cmder.exe` terminal and check if common Linux commands are supported.

```bash
ls --version
make --version
cp --version
chmod --version
rm --version
sed --version
```

<br>

**✅ Installing Sponge**

Open the `cmder.exe` terminal (not the built-in Windows cmd), and execute the following commands to install Sponge:

```bash
# Install Sponge
go install github.com/zhufuyi/sponge/cmd/sponge@latest

# Initialize Sponge, automatically install Sponge's dependency plugins
sponge init

# Check if all plugins have been successfully installed. If any plugins fail to install, retry with the command: sponge tools --install
sponge tools

# Check Sponge version
sponge -v
```

<br>

### Docker Deployment

> ⚠ Docker deployment is specifically for the Sponge UI service. If you need to develop based on the generated service code, you also need to install Sponge and the required plugins locally according to the installation instructions above.

**Option 1: Docker Run**

```bash
docker run -d --name sponge -p 24631:24631 zhufuyi/sponge:latest -l -a http://your_host_ip:24631
```

<br>

**Option 2: Docker Compose**

The content of the `docker-compose.yaml` file is as follows:

```yaml
version: "3.7"

services:
  sponge:
    image: zhufuyi/sponge:latest
    container_name: sponge
    restart: always
    command: ["-l","-a","http://your_host_ip:24631"]
    ports:
      - "24631:24631"
```

Start the service:

```bash
docker-compose up -d
```

After a successful Docker deployment, access `http://your_host_ip:24631` in your browser.
