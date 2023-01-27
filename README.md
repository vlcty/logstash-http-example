# logstash-http-example

I wanted to test how a custom plugin could be integrated into a logstash pipeline using the builtin http filter plugin. Two services written in python and go expose a HTTP API and return the received message from the body with a string prefix.   

## Usage

It's a docker compose project, so simply run:

> docker compose up

And you should see somethign like:

> logstash-http-plugin-logstash-1      | {   
> logstash-http-plugin-logstash-1      |                  "message" => "Sample line 1",   
> logstash-http-plugin-logstash-1      |     "http_response_python" => "Received: Sample line 1",   
> logstash-http-plugin-logstash-1      |         "http_response_go" => "Received: Sample line 1"   
> logstash-http-plugin-logstash-1      | }   
> logstash-http-plugin-logstash-1      | {   
> logstash-http-plugin-logstash-1      |                  "message" => "Sample line 2",   
> logstash-http-plugin-logstash-1      |     "http_response_python" => "Received: Sample line 2",   
> logstash-http-plugin-logstash-1      |         "http_response_go" => "Received: Sample line 2"   
> logstash-http-plugin-logstash-1      | }


The "message" field contains the string read from the `sample.log` file. It then get's sent to both services which prefix the received message with `Received: ` and return the result. Logstash then puts the respective answer either into the field `http_response_python` or `http_reponse_go`.

Exit the setup with Ctrl + C.