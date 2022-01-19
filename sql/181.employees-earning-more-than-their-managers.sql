--
-- @lc app=leetcode id=181 lang=mysql
--
-- [181] Employees Earning More Than Their Managers
--

-- @lc code=start
# Write your MySQL query statement below
select a.name as Employee
from Employee as a, Employee as b
where a.managerId = b.id
and a.salary > b.salary

-- @lc code=end

