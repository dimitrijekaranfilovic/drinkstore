export const usernameRules = [
  (username) => !!username || "Username is required!",
  (username) => username.length >= 3 || "Minimum username length must be 3!",
];
