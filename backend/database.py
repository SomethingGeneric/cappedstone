# standard library
import uuid
import os

# pypi
import toml

class EndpointDatabase:
    def __init__(self):
        self.base_path = "data/"
        if not os.path.exists(self.base_path):
            os.makedirs(self.base_path, exist_ok=True)

    def generate_endpoint_id(self):
        return uuid.uuid4()

    def register_endpoint(self, ip, hostname, osfamily, os, last_seen):
        obj = {
            "ip": ip,
            "hostname": hostname,
            "osfamily": osfamily,
            "os": os,
            "last_seen": last_seen,
            "next_expected": "",
            # TODO: next expected at (based on some config value for how often we want endpoints to be checking in)
        }

        id = self.generate_endpoint_id()

        with open(f"{self.base_path}/{id}.toml", "w") as f:
            f.write(toml.dumps(obj))


if __name__ == "__main__":
    e = EndpointDatabase()
    e.register_endpoint(
        input("IP: "),
        input("Hostname: "),
        input("OS Type: "),
        input("OS: "),
        input("Last seen: "),
    )
