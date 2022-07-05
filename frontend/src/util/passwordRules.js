export const passwordRules = [
  (password) => !!password || "Password is required!",
  (password) => password.length >= 6 || "Minimum password length is 6!",
];
