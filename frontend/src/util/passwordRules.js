export const passwordRules = [
  (password) => !!password || "Password is required!",
  (password) => password.length >= 6 || "Minimum password length is 6!",
  (password) =>
    /^(?=.*[0-9])(?=.*[!@#$%^&*?_=-\\+():\[\]])(?=.*[a-zA-Z])[a-zA-Z0-9!@#$%^&*?_=-\\+():\[\]]{6,}$/.test(
      password
    ) ||
    "Password must contain at least one number, one special character and one letter!",
];
