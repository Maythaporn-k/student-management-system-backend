import config from "../../../config.js";
import http from "k6/http";
import { check, sleep } from "k6";
import payload from "../../../payload.js";

const baseUrl = config.getOrchUrl();
const headers = config.headers();
let studentId = 0;
export function edit_user_orch() {
  studentId += 1;

  const Payload = payload.getEditPayload(studentId);

  const response = http.put(baseUrl + "/edit-user", Payload, { headers });

  // Add checks for the response
  const isSuccessful = check(response, {
    "status is 200": (r) => r.status === 200,
    "status is 409": (r) => r.status === 409,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });

  // Log error if the request fails
  if (response.status != 200) {
    console.error(
      `Request failed. Status: ${response.status}, Body: ${response.body} , ${Payload}`
    );
  } else {
    console.log(Payload);
  }

  sleep(1); // Simulate user wait time between requests
}
