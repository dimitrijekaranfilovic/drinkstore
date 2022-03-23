export const firstNameRules = [
  (username) => !!username || "First name is required!",
  (username) => username.length >= 2 || "Minimum first name length must be 2!",
];
