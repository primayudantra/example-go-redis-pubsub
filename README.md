**Publisher-Subscriber Example With Redis**

This repository contains a simple example of a publisher-subscriber pattern implemented in Go. The publisher sends messages, and the subscriber receives and processes them accordingly.

### Instructions:

1. **Clone Repository:**
   ```bash
   git clone https://github.com/primayudantra/example-go-redis-pubsub.git
   ```

2. **Navigate to Publisher Directory:**
   ```bash
   cd example-go-redis-pubsub/publisher
   ```

3. **Run Publisher:**
   ```bash
   go run publisher.go
   ```
   This command will start the publisher, which will begin sending messages.

4. **Open a New Terminal Window/Tab:**

5. **Navigate to Subscriber Directory:**
   ```bash
   cd ../subscriber
   ```

6. **Run Subscriber:**
   ```bash
   go run subscriber.go
   ```
   This command will start the subscriber, which will begin receiving and processing messages sent by the publisher.

### Additional Notes:

- The publisher and subscriber are implemented in separate Go files (`publisher.go` and `subscriber.go`) located in their respective directories.
- Ensure that Go is installed on your system and properly configured.
- The publisher and subscriber are designed to communicate locally. However, you can modify the code to support communication over a network if needed.
- Feel free to explore and modify the code to fit your requirements. This example serves as a basic demonstration of the publisher-subscriber pattern in Go.

### Requirements:

- Go programming language (Golang) installed on your system.
- Redis (Installation instruction [here](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/))
