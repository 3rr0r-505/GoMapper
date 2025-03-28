package cmd

// Probes is a list of payloads used for banner grabbing different services
var Probes = []string{
	// Generic
	"\n",
	"QUIT\r\n",
	"\x1b",

	// Web Services
	"HEAD / HTTP/1.0\r\n\r\n",
	"GET / HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"OPTIONS / HTTP/1.0\r\n\r\n",
	"TRACE / HTTP/1.0\r\n\r\n",
	"GET /server-status HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /manager/html HTTP/1.1\r\nHost: example.com\r\n\r\n",

	// Email Services
	"EHLO example.com\r\n",
	"HELO example.com\r\n",
	"VRFY root\r\n",
	"EXPN root\r\n",

	// FTP
	"USER anonymous\r\nPASS anonymous\r\n",
	"HELP\r\n",

	// SSH
	"\x53\x53\x48\x2D", // SSH protocol version banner request

	// SMB / Windows Services
	"\x80\x00\x02\x03",
	"\xfeSMB\r\n",
	"\x00\x83\x01\x00\x00\x00\x00\x00",

	// Remote Desktop (RDP)
	"\x03\x00\x00\x13\x0e\x00\x00\x00",

	// Databases
	"\x03\x00\x00\x00", // MySQL
	"\x04\x01\x00\x00", // PostgreSQL
	"Z",                // Oracle TNS
	"\x04\x02\x00\x00", // MSSQL
	"{ping:1}",         // MongoDB
	"\x00\x00\x00\x12\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", // Cassandra
	"\x40\x00\x00\x00", // CouchDB
	"INFO\r\n",         // Redis

	// Cloud & DevOps Services
	"GET /varz HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /metrics HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET / HTTP/1.1\r\nHost: etcd.example.com\r\n\r\n",
	"GET /health HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /api/v1/status HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /dashboard HTTP/1.1\r\nHost: example.com\r\n\r\n",

	// Security & Attack Frameworks
	"GET /ui HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /teamserver HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /status HTTP/1.1\r\nHost: example.com\r\n\r\n",
	"GET /admin HTTP/1.1\r\nHost: example.com\r\n\r\n",

	// Other Network Services
	"GET / HTTP/1.1\r\nHost: zookeeper.example.com\r\n\r\n",
	"GET / HTTP/1.1\r\nHost: consul.example.com\r\n\r\n",
	"\x00\x00\x00\x00\x00\x01\x00\x00stats\r\n",
	"OPTIONS sip:example.com SIP/2.0\r\n\r\n",
}
