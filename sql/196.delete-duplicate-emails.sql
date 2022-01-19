--
-- @lc app=leetcode id=196 lang=mysql
--
-- [196] Delete Duplicate Emails
--

-- @lc code=start
# Write your MySQL query statement below
delete p1 from Person p1, Person p2 
where p1.email = p2.email and p1.id > p2.id

-- @lc code=end

