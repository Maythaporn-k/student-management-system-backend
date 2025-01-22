import config from "../../../config.js";
import http from "k6/http";
import { check, sleep } from "k6";

const baseUrl = config.getOrchUrl();
const header = config.headers();

export function student_list_orch() {
  const response = http.get(baseUrl + "/student-list", {
    header,
  });

  // Add checks for the response
  const isSuccessful = check(response, {
    "status is 200": (r) => r.status === 200,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });

  // Log error if the request fails
  if (response.status != 200) {
    console.error(
      `Request failed. Status: ${response.status}, Body: ${response.body}`
    );
  } else {
    console.log("success");
  }

  sleep(1); // Simulate user wait time between requests
}
