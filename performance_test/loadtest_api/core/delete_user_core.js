import config from "../../../config.js";
import http from "k6/http";
import { check, sleep } from "k6";

const baseUrl = config.getCoreUrl();
const headers = config.headers();
let studentId = 0;

export function delete_user_core() {
  studentId += 1;
  const payload = JSON.stringify({ id: studentId });
  const response = http.del(baseUrl + "/delete-user", payload, { headers });

  // Add checks for the response
  const isSuccessful = check(response, {
    "status is 200": (r) => r.status === 200,
    "status is 409": (r) => r.status === 409,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });

  // Log error if the request fails
  if (r.status != 200) {
    console.error(
      `Request failed. Status: ${response.status}, Body: ${response.body}`
    );
  } else {
    console.log(Payload);
  }

  sleep(1); // Simulate user wait time between requests
}
