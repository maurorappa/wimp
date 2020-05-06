# wimp

What Is My Ip service designed to run in a container with low priviledges.

It reply with the IP present in a specified HTTP header for example the Haproxy 'X-Client-IP' or the CloudFlare 'CF-Connecting-IP'

this service provides two endpoints:

/s replies with jus the IP, ideal for scripting purposes due to its bare output

/d replies with more information taken from ipinfo service. You need to register and get an API key for this 
