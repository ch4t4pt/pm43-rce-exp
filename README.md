# Command Injection Exploit - PM43

This repository contains a command injection exploit for PM43 devices. The exploit targets the PM43 device with firmware version P10.11.013310 or earlier, which is vulnerable to command injection.

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/pm43-command-injection-exploit.git
   ```

2. Change into the project directory:

   ```bash
   cd pm43-command-injection-exploit
   ```

3. Build the exploit:

   ```bash
   go build
   ```

4. Run the exploit with the following command:

   ```bash
   ./pm43-command-injection-exploit -u <target-url> -p <injection-parameter> -c <linux-command>
   ```

   Replace `<target-url>` with the URL of the vulnerable PM43 device.

   Replace `<injection-parameter>` with the parameter to be injected. It should be either `username` or `userpassword`.

   Replace `<linux-command>` with the Linux command you want to execute on the target device.

5. The exploit will send a crafted POST request to the target device with the injected command. It will then check the response for successful command execution.

6. If the command is executed successfully, the output will be displayed on the console.

## Example

```bash
./pm43-command-injection-exploit -u http://target-device.com -p username -c "ls -la"
```

This example runs the exploit against the PM43 device at `http://target-device.com`, injecting the command `ls -la` into the `username` parameter.

## Disclaimer

This exploit is for educational and testing purposes only. Use it responsibly and with proper authorization. The author is not responsible for any misuse or damage caused by this exploit.

## Contributing

Contributions are welcome! If you find any issues or want to improve the exploit, feel free to create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).