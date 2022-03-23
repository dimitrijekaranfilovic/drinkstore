export const lastNameRules = [
  (username) => !!username || "Last name is required!",
  (username) => username.length >= 2 || "Minimum last name length must be 2!",
];
