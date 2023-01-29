# logstash-http-example

I wanted to test how a custom plugin could be integrated into a logstash pipeline using the builtin http filter plugin. A go plugin providing the HTTP service mimicks some transforming service.

**⚠️ This is a quick and dirty proof of concept example! Security or error handling is not a concern here!**

## Usage

It's a docker compose project, so simply run:

> docker compose up

And you should see something like:

> logstash-http-plugin-logstash-1  | {   
> logstash-http-plugin-logstash-1  |            "message" => "Sample fox 3",   
> logstash-http-plugin-logstash-1  |      "response_json" => {   
> logstash-http-plugin-logstash-1  |          "number_doubled" => 6,   
> logstash-http-plugin-logstash-1  |         "animal_repeated" => "foxfox"   
> logstash-http-plugin-logstash-1  |     },   
> logstash-http-plugin-logstash-1  |             "number" => "3",   
> logstash-http-plugin-logstash-1  |             "animal" => "fox",   
> logstash-http-plugin-logstash-1  |     "response_plain" => "Received: 3"   
> logstash-http-plugin-logstash-1  | }   
> logstash-http-plugin-logstash-1  | {   
> logstash-http-plugin-logstash-1  |            "message" => "Sample snake 4",   
> logstash-http-plugin-logstash-1  |      "response_json" => {   
> logstash-http-plugin-logstash-1  |          "number_doubled" => 8,   
> logstash-http-plugin-logstash-1  |         "animal_repeated" => "snakesnake"   
> logstash-http-plugin-logstash-1  |     },   
> logstash-http-plugin-logstash-1  |             "number" => "4",   
> logstash-http-plugin-logstash-1  |             "animal" => "snake",   
> logstash-http-plugin-logstash-1  |     "response_plain" => "Received: 4"   
> logstash-http-plugin-logstash-1  | }

Exit the setup with Ctrl + C.