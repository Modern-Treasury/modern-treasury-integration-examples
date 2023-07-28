## API example: Java + Spark

This small application demonstrates how you might set up a web server
using Java and [Spark][spark] with RESTful routes to accept your modern-treasury-js
form submissions without having to handle sensitive data.

This example makes use of the official [Java client library][client] for the Modern Treasury API.

Note that it is not necessary to use the Spark framework. In this example it is
used to organize various API actions into distinct application routes, but one
could just as easily implement these API actions within another application
framework.

### Use

#### Local

1. Install dependencies

The Modern Treasury Java library is distributed via [Maven Central](https://central.sonatype.com/artifact/com.moderntreasury/modern-treasury-java/1.1.0):

```xml
<dependency>
        <groupId>com.moderntreasury</groupId>
        <artifactId>modern-treasury-java</artifactId>
        <version>1.1.0</version>
</dependency>
```

The Spark Web Framework is distributed via [Maven Central](https://central.sonatype.com/artifact/com.sparkjava/spark-core/2.9.4)

```xml
<dependency>
    <groupId>com.sparkjava</groupId>
    <artifactId>spark-core</artifactId>
    <version>2.9.4</version>
</dependency>
```

2.  Start the server

```bash
mvn clean compile && mvn exec:java -Dexec.mainClass="com.moderntreasury.examples.App"
```

3. Open [http://localhost:9001](http://localhost:9001)

[spark]: https://sparkjava.com/
[client]: https://github.com/Modern-Treasury/modern-treasury-java