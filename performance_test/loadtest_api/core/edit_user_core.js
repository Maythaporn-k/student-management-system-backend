import config from "../../../config.js";
import http from "k6/http";
import { check, sleep } from "k6";
import payload from "../../../payload.js";

const baseUrl = config.getcoreUrl();
const headers = config.headers();
let studentId = 0;
export function edit_user_core() {
  studentId += 1;
  const Payload = payload.getEditPayload(studentId);
  const response = http.post(baseUrl + "/edit-user", Payload, { headers });

  // Add checks for the response
  const isSuccessful = check(response, {
    "status is 200": (r) => r.status === 200,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });

  // Log error if the request fails
  if (!isSuccessful) {
    console.error(
      `Request failed. Status: ${response.status}, Body: ${response.body}`
    );
  }

  sleep(1); // Simulate user wait time between requests
}
