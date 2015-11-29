# event-bridge

A little microdaemon that pulls messages out of journald and sends them to various log aggregators.

Supported destinations:

* Graylog, using GELF over UDP.
* Influx v0.9.

## Installing

If you have a go environment configured, you can install it with:

    go get github.com/altaireltd/event-bridge

Alternatively, you can use a binary from https://github.com/altaireltd/event-bridge/releases

## Running

It can be run manually like

    GRAYLOG_ADDRESS=1.2.3.4:12201 INFLUX_URL=http://1.2.3.4:8086/ event-bridge
