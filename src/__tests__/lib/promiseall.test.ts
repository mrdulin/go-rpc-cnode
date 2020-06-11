import { promiseAll } from "../../lib/promiseAll";

describe("lib", () => {
  describe("promiseAll", () => {
    it("should pass", async () => {
      const promises = Array.from({ length: 10 }).map((_, i) =>
        Promise.resolve(i)
      );
      const rs = await promiseAll(promises);
      expect(rs).toEqual([0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);
    });

    it("should reject all if any promise reject ", async () => {
      const promises = [
        Promise.resolve(1),
        Promise.reject(new Error("network")),
      ];
      await expect(promiseAll(promises)).rejects.toThrowError("network");
    });
  });
});
