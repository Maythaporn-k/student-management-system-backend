const name = ["Jack ", "Susan ", "Thomas ", "Ali ", "Florence "][
  Math.floor(Math.random() * 5)
];
const surName = ["Dawson", "Storm", "Elva", "Abhams", "Shinya"][
  Math.floor(Math.random() * 5)
];
export default class payload {
  static getCreatePayload() {
    const payload = JSON.stringify({
      name: name + surName,
      age: Math.floor(Math.random() * (18 - 10 + 1)) + 10,
      grade: ["A", "B", "C", "D"][Math.floor(Math.random() * 4)],
    });

    return payload;
  }

  static getEditPayload(studentId = 1) {
    const payload = JSON.stringify({
      id: studentId,
      name: name + surName,
      age: Math.floor(Math.random() * (18 - 10 + 1)) + 10, // Random age between 10 and 18
      grade: ["A", "B", "C", "D", "F"][Math.floor(Math.random() * 5)], // Random grade
      attendance: [true, false][Math.floor(Math.random() * 2)],
    });

    return payload;
  }
}
