use std::io::Write;
use std::net::TcpStream;
use std::process;

fn main() {
    // Get the hostname
    let hostname = match hostname::get() {
        Ok(name) => name.to_string_lossy().to_string(),
        Err(e) => {
            eprintln!("Failed to get hostname: {}", e);
            process::exit(1);
        }
    };

    println!("Connecting to localhost:65432...");
    
    // Connect to the TCP server
    let mut stream = match TcpStream::connect("localhost:65432") {
        Ok(stream) => {
            println!("Successfully connected to localhost:65432");
            stream
        },
        Err(e) => {
            eprintln!("Failed to connect to localhost:65432: {}", e);
            process::exit(1);
        }
    };

    // Send the hostname
    println!("Sending hostname: {}", hostname);
    if let Err(e) = stream.write_all(hostname.as_bytes()) {
        eprintln!("Failed to send hostname: {}", e);
        process::exit(1);
    }

    // Send a newline to ensure the server receives the complete message
    if let Err(e) = stream.write_all(b"\n") {
        eprintln!("Failed to send newline: {}", e);
        process::exit(1);
    }

    // Flush the stream to ensure data is sent
    if let Err(e) = stream.flush() {
        eprintln!("Failed to flush stream: {}", e);
        process::exit(1);
    }

    println!("Hostname sent successfully. Exiting.");
}