export default class config {
  static headers() {
    return {
      "content-type": "application/json",
    };
  }

  static getOrchUrl() {
    return "http://localhost:3001/orch";
  }

  static getCoreUrl() {
    return "http://localhost:3002/core";
  }
}
