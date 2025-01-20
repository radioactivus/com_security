# Security Project

## Project Description: Remote Secure File Transfer between Linux Devices

This project involves building a secure, client-server file transfer system between two Linux devices, with a strong focus on modern Go programming, concurrency, and Linux system knowledge. The system will utilize encryption algorithms, certificate management, and system-level automation to ensure a secure, efficient, and user-specific file transfer process.

---

## Key Features

### Client/Server Architecture
- The server runs a file transfer service on one Linux device.
- The client sends commands to securely transfer files to/from the server.

### Systemd Integration
- Use systemd socket activation for communication.
- Maintain logs using systemd's journal for debugging and audit trails.

### Secure Sending Command
- Provide a command-line tool for users to securely send files.
- Ensure all communication between client and server is encrypted.

### Predefined Path
- Restrict file transfers to a predefined directory on the server (e.g., `/home/user/secure-ftps/`).

### User-Specific Access
- Authenticate users and restrict them to their own directories within the predefined path.
- Limit file sizes or total storage for each user.

---

## Implementation Details

### 1. Communication
- Use systemd sockets to listen for incoming client connections.
- Ensure that all communication uses TLS/SSL for encryption.
- Use Go's `crypto/tls` package for secure communication.

### 2. File Transfer
- Encrypt files using symmetric encryption (e.g., AES-256) before transfer.
- Use asymmetric encryption (e.g., RSA) for secure key exchange.

### 3. User Authentication
- Authenticate users with credentials stored securely (e.g., hashed passwords in `/etc`).
- Use PAM (Pluggable Authentication Modules) or create a lightweight authentication system in Go.

### 4. Secure Commands
- Implement a CLI tool (e.g., `secure-send`) in Go that:
  - Encrypts the file before sending.
  - Sends the file to the server over the socket.
  - Verifies the integrity of the file after transfer.

### 5. Logging
- Log all file transfer attempts and results via systemd's journal.
- Include details like timestamp, user ID, file path, and transfer success/failure.

### 6. Predefined Path
- Enforce that all files sent to the server are stored in a specific directory.
- Use Linux permissions or chroot environments to isolate user data.

### 7. Benchmarking Implementation
#### Encryption Algorithm Testing
- Implement multiple algorithms using Go's `crypto` package.
- Measure and compare:
  - Encryption/Decryption time.
  - Resource usage (CPU, memory).
#### Chunk Size Testing
- Implement file chunking in the transfer process.
- Dynamically adjust chunk sizes and measure their impact on transfer speed and resource usage.
#### Command-line Options
- Add options to run benchmarks (e.g., `--benchmark-algorithms`, `--benchmark-chunks`).

---

## Certificate Management

### Automated Certificate Generation
- Include a script (Python or Bash) to create self-signed X.509 certificates.
- Use appropriate extensions for server and client roles (e.g., `serverAuth` and `clientAuth`).

### Secure Exchange
- Implement a mechanism to exchange and verify certificates during the handshake.
- Store certificates locally in a secure path (e.g., `/etc/secure-transfer/certs/`).

### Revocation and Renewal
- Add a process for certificate expiration checks and automated renewal.

---

## System Requirements

### Hardware
- Two Linux devices with sufficient CPU and RAM.
- Minimum 512MB RAM, 1GB recommended.
- At least 100MB of available storage for dependencies, logs, and files.

### Operating System
- A Linux-based OS such as Debian, Ubuntu, or similar distributions.
- Kernel version 4.9 or higher to ensure compatibility with Go and OpenSSL libraries.

### Software Dependencies
- Go 1.18 or higher for building and running the application.
- OpenSSL 1.1+ for TLS and encryption (where needed).
- Systemd (default in most Linux distributions).
- Python 3.x or Bash for automation scripts.

### Network Requirements
- Secure SSH or TCP connections between devices.
- Sufficient bandwidth for file transfers.

---

## Linux-Specific Requirements

### Command-line Tools
- Utilize and integrate Linux command-line tools such as:
  - `openssl`: For certificate creation and encryption testing.
  - `scp` / `rsync`: To benchmark against custom implementation.
  - `journalctl`: To analyze systemd logs for debugging.
  - `stat`, `ls`, `chmod`, `chown`: For managing files, permissions, and ownership in predefined paths.
  - `curl` or `wget`: For testing HTTP/HTTPS endpoints if implemented.

### Systemd Integration
- **Socket Activation:**
  - Use systemd sockets to manage client-server communication.
  - Demonstrate understanding of socket-based activation by configuring and testing `.socket` and `.service` files.
- **Automating Startup:**
  - Create and deploy a systemd unit file for the server to start automatically at boot.
  - Include environment variable configurations for customization.
- **Logging:**
  - Ensure all logs are written to systemd’s journal and accessible via `journalctl`.

### Application Building
- Use Go modules to manage dependencies and build the project, emphasizing:
  - Modular design for Go code (e.g., separate packages for encryption, communication, etc.).
  - Dependency management (e.g., linking Go's `crypto` library).
  - Debug and release builds with optimizations.
- Package the application as a `.deb` file for easy installation, including systemd unit files.

### Discretionary Access Control (DAC)
- Implement strict file permissions to secure sensitive files and directories:
  - Ensure only specific users can access transfer-related directories.
  - Use `chmod` and `chown` to enforce access restrictions.
- Explain the significance of DAC in the project documentation.
- Optionally simulate a scenario where DAC prevents unauthorized access to server files.

