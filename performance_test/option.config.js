export default class Options {
  static const(
    thresholds = null //default
  ) {
    const options = {
      vus: 10, // จำนวนผู้ใช้งานเริ่มต้น
      duration: "60s", // ระยะเวลารันการทดสอบ
    };

    if (thresholds != null) {
      options.thresholds = thresholds;
    }

    return options;
  }

  static const_arrive_time(thresholds = null) {
    const options = {
      scenarios: {
        student_loadtest: {
          executor: "constant-arrival-rate",
          rate: 5,
          timeUnit: "10s",
          duration: "30s",
          preAllocatedVUs: 20,
        },
      },
    };

    // หาก thresholds ไม่เป็น null ให้เพิ่มเข้าไปใน options
    if (thresholds != null) {
      options.thresholds = thresholds;
    }

    return options; // คืนค่ากลับ
  }
}
