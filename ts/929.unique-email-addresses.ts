/*
 * @lc app=leetcode id=929 lang=typescript
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
function numUniqueEmails(emails: string[]): number {
  const uniques = new Set<string>();

  for (let email of emails) {
    let [local, domain] = email.split("@");
    [local] = local.split("+");

    const validEmail = `${local.replace(/\./g, "")}@${domain}`;

    uniques.add(validEmail);
  }

  return uniques.size;
}
// @lc code=end
