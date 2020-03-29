import http from "k6/http";
import { check } from "k6";

export let options = {
    stages: [
        { duration: '20s', target: 10 },
        { duration: '15s', target: 60 },
        { duration: '20s', target: 220 },
    ],
    thresholds: { http_req_duration: ['avg<100', 'p(95)<1000'] },
    userAgent: 'loadtest/1.0',
};

export default function() {
    // Send a JSON encoded POST request
    let body = JSON.stringify({
        email: "test@login.com",
        password: "mysecretpassword",
    });
    let response = http.post("http://localhost:8080/api/v1/login", body, { headers: { "Content-Type": "application/json" }});

    // Use JSON.parse to deserialize the JSON (instead of using the r.json() method)
    let data = JSON.parse(response.body);

    // Verify response
    check(response, {
        "status is 200": (r) => r.status === 200,
        "is key correct": (r) => data.token !== "",
    });
}
