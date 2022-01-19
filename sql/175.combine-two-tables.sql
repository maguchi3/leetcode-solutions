--
-- @lc app=leetcode id=175 lang=mysql
--
-- [175] Combine Two Tables
--

-- @lc code=start
# Write your MySQL query statement below

select firstName, lastName, city, state
from Person left join Address
on Person.personId = Address.personId

-- @lc code=end

