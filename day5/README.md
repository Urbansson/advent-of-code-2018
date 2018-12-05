


## --- Day 5: Alchemical Reduction ---

You&#39;ve managed to sneak in to the prototype suit manufacturing lab.  The Elves are making decent progress, but are still struggling with the suit&#39;s size reduction capabilities.

While the very latest in 1518 alchemical technology might have solved their problem eventually, you can do better.  You scan the chemical composition of the suit&#39;s material and discover that it is formed by extremely long [polymers](https://en.wikipedia.org/wiki/Polymer) (one of which is 
available
 as your puzzle input).

The polymer is formed by smaller _units_ which, when triggered, react with each other such that two adjacent units of the same type and opposite polarity are destroyed. Units&#39; types are represented by letters; units&#39; polarity is represented by capitalization.  For instance, `r` and `R` are units with the same type but opposite polarity, whereas `r` and `s` are entirely different types and do not react.

For example:

*   In `aA`, `a` and `A` react, leaving nothing behind.
*   In `abBA`, `bB` destroys itself, leaving `aA`.  As above, this then destroys itself, leaving nothing.
*   In `abAB`, no two adjacent units are of the same type, and so nothing happens.
*   In `aabAAB`, even though `aa` and `AA` are of the same type, their polarities match, and so nothing happens.

Now, consider a larger example, `dabAcCaCBAcCcaDA`:

``dabA_cC_aCBAcCcaDA  The first &#39;cC&#39; is removed.
dab_Aa_CBAcCcaDA    This creates &#39;Aa&#39;, which is removed.
dabCBA_cCc_aDA      Either &#39;cC&#39; or &#39;Cc&#39; are removed (the result is the same).
dabCBAcaDA        No further actions can be taken.
``

After all possible reactions, the resulting polymer contains _10 units_.

_How many units remain after fully reacting the polymer you scanned?_ 
(Note: in this puzzle and others, the input is large; if you copy/paste your input, make sure you get the whole thing.)






## --- Part Two ---

Time to improve the polymer.

One of the unit types is causing problems; it&#39;s preventing the polymer from collapsing as much as it should.  Your goal is to figure out which unit type is causing the most problems, remove all instances of it (regardless of polarity), fully react the remaining polymer, and measure its length.

For example, again using the polymer `dabAcCaCBAcCcaDA` from above:

*   Removing all `A`/`a` units produces `dbcCCBcCcD`. Fully reacting this polymer produces `dbCBcD`, which has length 6.
*   Removing all `B`/`b` units produces `daAcCaCAcCcaDA`. Fully reacting this polymer produces `daCAcaDA`, which has length 8.
*   Removing all `C`/`c` units produces `dabAaBAaDA`. Fully reacting this polymer produces `daDA`, which has length 4.
*   Removing all `D`/`d` units produces `abAcCaCBAcCcaA`. Fully reacting this polymer produces `abCBAc`, which has length 6.

In this example, removing all `C`/`c` units was best, producing the answer _4_.

_What is the length of the shortest polymer you can produce_ by removing all units of exactly one type and fully reacting the result?


