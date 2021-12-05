package main

import (
  "fmt"
  "io/ioutil"
  "math"
  "regexp"
  // "sort"
  "strings"
  "strconv"
  "unicode/utf8"
)

// func AoCDayEleventeen() {
//   input, _ := ioutil.ReadFile("day_111_input.txt")
//   lines := strings.Split(string(input), "\n")
//   result_part_one := DayEleventeenPartOne(lines)
//   fmt.Printf("111.1: %v\n", result_part_one)
//   result_part_two := DayEleventeenPartTwo(lines)
//   fmt.Printf("111.2: %v\n", result_part_two)
// }

// func DayEleventeenPartOne(data []string) int {
//   return -1
// }

// func DayEleventeenPartTwo(data []string) int {
//   return -1
// }

func main() {
  // AoCDayOne()
  // AoCDayTwo()
  // AoCDayThree()
  // AoCDayFour()
  // AoCDayFive()
  // AoCDaySix()
  // AoCDaySeven()
  // AoCDayEight()
  AoCDayNine()
}

func AoCDayNine() {
  input, _ := ioutil.ReadFile("day_9_input.txt")
  lines := strings.Split(string(input), "\n")
  result_part_one := DayNinePartOne(lines)
  fmt.Printf("9.1: %v\n", result_part_one)
  result_part_two := DayNinePartTwo(lines)
  fmt.Printf("9.2: %v\n", result_part_two)
}

func DayNinePartOne(data []string) int {
  return -1
}

func DayNinePartTwo(data []string) int {
  return -1
}

func AoCDayEight() {
  input, _ := ioutil.ReadFile("day_8_input.txt")
  lines := strings.Split(string(input), "\n")
  result_part_one := DayEightPartOne(lines)
  result_part_two := DayEightPartTwo(lines)
  fmt.Printf("8.1: %v\n8.2: %v\n", result_part_one, result_part_two)
}

type Index int

type Instruction struct {
  ix  Index
  op  string
  val int
}

func ParseInstructions(input []string) (instructions []Instruction) {
  instructions = make([]Instruction, 0, len(input))
  for ix, command_str := range input {
    command := strings.Split(command_str, " ")
    val, err := strconv.Atoi(command[1])
    if err != nil {
      panic("couldn't parse instruction")
    }
    instructions = append(instructions, Instruction{
      ix: Index(ix), 
      op: command[0], 
      val: val,
    })
  }
  return
}

func DayEightPartOne(input []string) (acc int) {
  instructions := ParseInstructions(input)
  seen_ix := make(map[Index]struct{})
  cur_ix := Index(0)
  for {
    if _, found := seen_ix[cur_ix]; found {
      break
    }
    seen_ix[cur_ix] = struct{}{}
    cur_instruction := instructions[cur_ix]
    switch cur_instruction.op {
      case "jmp":
      cur_ix = Index(int(cur_ix) + cur_instruction.val)
      case "acc": 
      acc += cur_instruction.val
      fallthrough
      case "nop":
      cur_ix += 1
    }
  }

  return
}

type SavedState struct {
  instruction Instruction
  acc         int
}

func DayEightPartTwo(input []string) (acc int) {
  instructions := ParseInstructions(input)
  seen_ix := make(map[Index]struct{}, len(input))
  swap_candidates := make([]SavedState, 0, len(input)/2)
  cur_ix := Index(0)
  for {
    if _, found := seen_ix[cur_ix]; found {
      break
    }
    seen_ix[cur_ix] = struct{}{}
    cur_instruction := instructions[cur_ix]
    switch cur_instruction.op {
      case "jmp":
      swap_candidates = append(swap_candidates, SavedState{cur_instruction, acc})
      cur_ix = Index(int(cur_ix) + cur_instruction.val)
      case "acc":
      acc += cur_instruction.val
      fallthrough
      case "nop":
      cur_ix += 1
    }
  }

  OUTER:
  for _, candidate := range swap_candidates {
    acc = candidate.acc
    cur_ix := candidate.instruction.ix
    delete(seen_ix, cur_ix)
    cur_op := instructions[cur_ix].op
    if cur_op == "jmp" {
      instructions[cur_ix].op = "nop"
    } else {
      instructions[cur_ix].op = "jmp"
    }

    for {
      if int(cur_ix) >= len(instructions) {
        // fmt.Println("SUCCESS")
        break OUTER
      }
      if _, found := seen_ix[cur_ix]; found {
        break
      }

      cur_instruction := instructions[cur_ix]
      seen_ix[cur_ix] = struct{}{}

      // // DEBUG -- START
      // fmt.Printf(":%d Running ", cur_ix)
      // switch cur_instruction.op {
      //   case "jmp":
      //   fmt.Printf("jmp %d, :%d -> :%d\n", cur_instruction.val, int(cur_ix), int(cur_ix) + cur_instruction.val)
      //   case "nop":
      //   fmt.Printf("nop %d\n", cur_instruction.val)
      //   case "acc":
      //   fmt.Printf("acc %d (%d -> %d)\n", cur_instruction.val, acc, acc + cur_instruction.val)
      // }
      // fmt.Printf("    before acc: %d @ :%d\n", acc, int(cur_ix))
      // // DEBUG -- END

      switch cur_instruction.op {
        case "jmp":
        cur_ix = Index(int(cur_ix) + cur_instruction.val)
        case "nop":
        cur_ix += 1
        case "acc":
        acc += cur_instruction.val
        cur_ix += 1
      }
      // fmt.Printf("    after acc: %d @ :%d\n", acc, int(cur_ix))
    }
  }

  return


  // instructions := ParseInstructions(input)
  // seen_ix := make(map[Index]struct{}, len(input))
  // jmp_nop_stack := make([]SavedState, 0, len(input)/2)
  // cur_ix := Index(0)
  // for {
  //   if int(cur_ix) >= len(instructions) {
  //     break
  //   }

  //   cur_instruction := instructions[cur_ix]
  //   fmt.Printf(":%d Running ", cur_ix)
  //   switch cur_instruction.op { 
  //     case "jmp":
  //     fmt.Printf("jmp %d, :%d -> :%d\n", cur_instruction.val, int(cur_ix), int(cur_ix) + cur_instruction.val)
  //     case "nop": 
  //     fmt.Printf("nop %d\n", cur_instruction.val)
  //     case "acc": 
  //     fmt.Printf("acc %d (%d -> %d)\n", cur_instruction.val, acc, acc + cur_instruction.val)
  //   }
  //   fmt.Printf("    before acc: %d @ :%d\n", acc, int(cur_ix))

  //   is_loaded_state := false
  //   if _, found := seen_ix[cur_ix]; found {
  //     // Found a loop - restart from saved state
  //     is_loaded_state = true
  //     last_state_ix := len(jmp_nop_stack) - 1
  //     popped_state := jmp_nop_stack[last_state_ix]
  //     jmp_nop_stack = jmp_nop_stack[:last_state_ix]

  //     // Load state
  //     cur_ix = popped_state.instruction.ix
  //     val := popped_state.instruction.val
  //     acc = popped_state.acc
  //     fmt.Printf(" Reverting %s %d, acc: %d @ :%d\n", popped_state.instruction.op, val, acc, popped_state.instruction.ix)

  //     // Swap JMP and NOP
  //     next_op := "jmp"
  //     if popped_state.instruction.op == "jmp" {
  //       next_op = "nop"
  //     }
  //     cur_instruction = Instruction{
  //       ix: cur_ix, 
  //       op: next_op,
  //       val: val,
  //     }

  //     fmt.Printf(" running ")
  //     switch cur_instruction.op {
  //       case "jmp":
  //       fmt.Printf("jmp %d (%d -> %d)\n", cur_instruction.val, int(cur_ix), int(cur_ix) + cur_instruction.val)
  //       case "nop": 
  //       fmt.Printf("nop %d\n", cur_instruction.val)
  //     }
  //     fmt.Printf("    restored acc: %d @ :%d\n", acc, int(cur_ix))
  //   }

  //   seen_ix[cur_ix] = struct{}{}

  //   switch cur_instruction.op {
  //     case "jmp": 
  //     if !is_loaded_state {
  //       fmt.Printf("  saving jmp %d, acc: %d @ :%d\n", cur_instruction.val, acc, int(cur_ix))
  //       jmp_nop_stack = append(jmp_nop_stack, SavedState{cur_instruction, acc})
  //     }
  //     cur_ix = Index(int(cur_ix) + cur_instruction.val)
  //     case "nop": 
  //     if !is_loaded_state {
  //       fmt.Printf("  saving nop %d, acc: %d @ :%d\n", cur_instruction.val, acc, int(cur_ix))
  //       jmp_nop_stack = append(jmp_nop_stack, SavedState{cur_instruction, acc})
  //     }
  //     cur_ix += 1
  //     case "acc": 
  //     acc += cur_instruction.val
  //     cur_ix += 1
  //   }
  //   fmt.Printf("    after acc: %d @ :%d\n", acc, int(cur_ix))
  // }

  // return
}

func AoCDaySeven() {
  input, _ := ioutil.ReadFile("day_7_input.txt")
  lines := strings.Split(string(input), "\n")
  result := DaySevenPartOne(lines)
  fmt.Println("7.1:", result)
  result_part_two := DaySevenPartTwo(lines)
  fmt.Println("7.2:", result_part_two)
}

func MakeRuleMap(rules []string) (rule_map map[string]map[string]uint16) {
  rule_map = make(map[string]map[string]uint16)
  rule_regex := regexp.MustCompile(
    `((?:\w+ ){1,2})bags contain ((?:(?:\d+ (?:\w+ ){1,2}|no other )bags?[,.]? ?)+)`)
  contained_bags_regex := regexp.MustCompile(
    `(\d+) ((?:\w+ ){1,2})bags?[,.]`)
  for _, rule_str := range rules {
    rule := rule_regex.FindStringSubmatch(rule_str)
    bag := strings.Trim(rule[1], " ")
    rule_map[bag] = make(map[string]uint16)

    contained_bags_results := contained_bags_regex.FindAllStringSubmatch(rule[2], -1)
    for _, contained_bag_result := range contained_bags_results {
      contained_quantity, err := strconv.ParseUint(contained_bag_result[1], 10, 16)
      contained_bag_contents := strings.Trim(contained_bag_result[2], " ")
      if err != nil {
        panic("failed to parse contained bag")
      }
      rule_map[bag][contained_bag_contents] = uint16(contained_quantity)
    }
  }
  // for ka, va := range rule_map { 
  //   // if len(va) == 0 {
  //   //   fmt.Println(ka)
  //   // }
  //   // fmt.Printf("%s\n  ", ka)
  //   for kb, vb := range va {
  //     _ = ka
  //     _ = kb
  //     _ = vb
  //     // fmt.Printf("%d %s  ", vb, kb)
  //   }
  //   // fmt.Println()
  // }

  return
}


func DaySevenPartOne(data []string) (num_bags int) {
  contains_map := MakeRuleMap(data)
  contained_map := make(map[string]map[string]struct{})
  for container_bag, contained_bags := range contains_map {
    for contained_bag, _ := range contained_bags {
      if _, found := contained_map[contained_bag]; !found {
        contained_map[contained_bag] = make(map[string]struct{})
      }
      contained_map[contained_bag][container_bag] = struct{}{}
    }
  }

  // if _, found := rule_map["shiny gold"]; found {
    // fmt.Println("found gold")
  // }

  // for contained_bag, container_bag_set := range contained_map {
  //   fmt.Printf("contained: %s\n  containers:", contained_bag)
  //   for container_bag := range container_bag_set {
  //     fmt.Printf(" %s", container_bag)
  //   }
  //   fmt.Println()
  // }

  // fmt.Println(contained_map["clear salmon"])
  // fmt.Println(contains_map["clear silver"])

  seen_bags := make(map[string]struct{})
  init_container_bags := contained_map["shiny gold"]
  check_bags := make([]string, 0)
  for container_bag := range init_container_bags {
    check_bags = append(check_bags, container_bag)
    // fmt.Println("Bag queued: ", container_bag)
  }

  for len(check_bags) != 0 {
    bag := check_bags[0]
    check_bags = check_bags[1:]
    seen_bags[bag] = struct{}{}
    // fmt.Println("Bag dequeued: ", bag)
    container_bags := contained_map[bag]
    for container_bag := range container_bags {
      // fmt.Println("Bag queued: ", container_bag)
      check_bags = append(check_bags, container_bag)
    }
  }

  num_bags = len(seen_bags)


  // NOTE: This is a CSP attempt. However, it deadlocks
  //   because there's no real way to tell when we're done 
  //   checking the channel... Interesting concurrency problem.

  // check_bags := make(chan string, len(contained_map))
  // init_container_bags := contained_map["shiny gold"]
  // for container_bag := range init_container_bags {
  //   fmt.Println("Check bag: ", container_bag)
  //   check_bags <- container_bag 
  // }

  // for bag := range check_bags {
  //   seen_bags[bag] = struct{}{}
  //   fmt.Println("Received bag: ", bag)
  //   container_bags := contained_map[bag]
  //   for container_bag := range container_bags {
  //     fmt.Println("Check bag: ", bag)
  //     check_bags <- container_bag
  //   }
  // }

  return
}

func DaySevenPartTwo(data []string) (bag_sum uint) {
  contains_map := MakeRuleMap(data)
  bag_sums := make(map[string]uint)

  var getBagSum func(bag string) (bag_sum uint)
  getBagSum = func (bag string) (bag_sum uint) {
    if cached_bag_sum, found := bag_sums[bag]; found {
      // fmt.Printf("  Found cached bag %s\n", bag)
      bag_sum = cached_bag_sum
    } else {
      // fmt.Printf("%s contains\n", bag)
      // for contained_bag, quantity := range contains_map[bag] {
      //   fmt.Printf("  %d %s\n", quantity, contained_bag)
      // }
      // fmt.Println("\n")

      bag_sum = 0
      for contained_bag, quantity := range contains_map[bag] {
        contained_bag_sum := getBagSum(contained_bag)
        // fmt.Printf("  %s bag sum: %d\n", contained_bag, contained_bag_sum)
        bag_sum += contained_bag_sum * uint(quantity) + uint(quantity)
      }
      bag_sums[bag] = bag_sum
    }
    return 
  }

  bag_sum = getBagSum("shiny gold")
  // fmt.Println(bag_sums)

  return
}

func AoCDaySix() {
  separator := "\n\n"
  input, _ := ioutil.ReadFile("day_6_input.txt")
  lines := strings.Split(string(input), separator)
  result := DaySixPartOne(lines)
  fmt.Println("6.1:", result)
  result = DaySixPartTwo(lines)
  fmt.Println("6.2:", result)
}

func DaySixPartOne(data []string) (yes_count int) {
  for _, line := range data {
    seen := make(map[rune]struct{})
    for _, q := range line {
      if q != '\n' {
        // fmt.Println(string(q))
        seen[q] = struct{}{}
      }
    }
    yes_count += len(seen)
  }
  return yes_count
}

func DaySixPartTwo(data []string) (group_yes_count int) {
  for _, line := range data {
    answers := strings.Split(line, "\n")
    group_size := len(answers)

    // Count number of yeses for each question
    seen := make(map[rune]int)
    for _, answer := range answers {
      for _, q := range answer {
        seen[q] += 1
        // fmt.Println(answer, string(q), seen[q])
      }
    }

    // Count group totals
    for _, yes_count := range seen {
      if yes_count == group_size {
        group_yes_count += 1
      }
    }
  }

  return
}

func AoCDayFive() {
  input, _ := ioutil.ReadFile("day_5_input.txt")
  lines := strings.Split(string(input), "\n")
  result := DayFivePartOne(lines)
  fmt.Println("5.1:", result)
  result = DayFivePartTwo(lines)
  fmt.Println("5.2:", result)
}

func SeatIdFromRowCol(row uint16, col uint16) uint16 {
  return (row * 8) + col
}

func SeatId(pass string) uint16 {
  row_str := pass[:7]
  col_str := pass[7:]
  // fmt.Println(row_str, col_str)

  // convert to binary
  row_bin_str := make([]rune, 0, len(row_str))
  col_bin_str := make([]rune, 0, len(col_str))
  for _, c := range row_str {
    if c == 'B' {
      row_bin_str = append(row_bin_str, '1')
    } else {
      row_bin_str = append(row_bin_str, '0')
    }
  }
  for _, c := range col_str {
    if c == 'R' {
      col_bin_str = append(col_bin_str, '1')
    } else {
      col_bin_str = append(col_bin_str, '0')
    }
  }
  // fmt.Println(row_bin_str, col_bin_str)

  row, err := strconv.ParseUint(string(row_bin_str), 2, 16)
  if err != nil {
    return 0
  }
  col, err := strconv.ParseUint(string(col_bin_str), 2, 16)
  if err != nil {
    return 0
  }
  // fmt.Println(row, col)

  return SeatIdFromRowCol(uint16(row), uint16(col))
}

func DayFivePartOne(passes []string) (highest_seat_id uint16) {
  for _, pass := range passes {
    seat_id := SeatId(pass)    
    if seat_id > highest_seat_id {
      highest_seat_id = seat_id
    }
  }

  return
}

func DayFivePartTwo(passes []string) (empty_seat_id uint16) {
  num_seats := uint16(math.Exp2(10))
  filled_seats := make(map[uint16]bool, int(num_seats))
  for _, pass := range passes {
    filled_seats[SeatId(pass)] = true
  }

  start_search := false
  for seat_id := uint16(0); seat_id < num_seats; seat_id++ {
    // fmt.Println("Found", seat_id, num_seats, filled_seats[seat_id])
    if start_search && filled_seats[seat_id] != true {
      empty_seat_id = seat_id
      break
    }
    start_search = filled_seats[seat_id]
  }

  return
}

func AoCDayFour() {
  separator := "\n\n"
  input, _ := ioutil.ReadFile("day_4_input.txt")
  lines := strings.Split(string(input), separator)
  result := DayFourPartOne(lines)
  fmt.Println("4.1:", result)
  result = DayFourPartTwo(lines)
  fmt.Println("4.2:", result)
}

func DayFourPartOne(passports []string) (valid_count int) {
  searched_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
  passport_regex_str := fmt.Sprintf(`(%s):[^\s]+`, strings.Join(searched_fields, "|"))
  passport_regex := regexp.MustCompile(passport_regex_str)
  for _, passport := range passports {
    matches := passport_regex.FindAllStringSubmatch(passport, -1)
    if len(matches) >= len(searched_fields) { 
      valid_count += 1
    }
  }
  return
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
//   If cm, the number must be at least 150 and at most 193.
//   If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func DayFourPartTwo(passports []string) (valid_count int) {
  searched_fields_to_regex := map[string]*regexp.Regexp{
    // I hate this. Forgot the ^ and $...
    "byr": regexp.MustCompile(`^19[2-9][0-9]$|^200[0-2]$`), 
    "iyr": regexp.MustCompile(`^20(1[0-9]|20)$`), 
    "eyr": regexp.MustCompile(`^20(2[0-9]|30)$`), 
    "hgt": regexp.MustCompile(`^1([5-8][0-9]|9[0-3])cm$|^(59|6[0-9]|7[0-6])in`), 
    "hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`), 
    "ecl": regexp.MustCompile(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`), 
    "pid": regexp.MustCompile(`^\d{9}$`),
  }
  searched_fields := make([]string, 0, len(searched_fields_to_regex))
  for key := range searched_fields_to_regex {
    searched_fields = append(searched_fields, key)
  }
  passport_regex_str := fmt.Sprintf(`(%s):([^\s]+)`, strings.Join(searched_fields, `|`))
  passport_regex := regexp.MustCompile(passport_regex_str)

  // // DEBUG -- START
  // distinct_vals := make(map[string]map[string]struct{})
  // // DEBUG -- END

  OUTER:
  for _, passport := range passports {
    matches := passport_regex.FindAllStringSubmatch(passport, -1)

    // continue if invalid
    if len(matches) != len(searched_fields) { 
      continue OUTER
    } else {
      for _, match := range matches {
        field := match[1]
        value := match[2]
        field_regex, found := searched_fields_to_regex[field]
        if !found || !field_regex.MatchString(value) {
          continue OUTER
        }
        // // DEBUG -- START
        // if field_vals, found := distinct_vals[field]; found {
        //   field_vals[value] = struct{}{}
        // } else {
        //   distinct_vals[field] = map[string]struct{}{ value: struct{}{} }
        // }
        // // DEBUG -- END
      }
    }

    valid_count += 1
  }

  // // DEBUG -- START
  // for field, field_vals := range distinct_vals { 
  //   values := []string{}
  //   for val := range field_vals {
  //     values = append(values, val)
  //   }
  //   sort.Strings(values)
  //   _ = field
  //   // fmt.Printf("%s:\n%s\n", field, strings.Join(values, "  \n"))
  // }
  // // DEBUG -- END
  
  return
}

func AoCDayThree() {
  input, _ := ioutil.ReadFile("day_3_input.txt")
  lines := strings.Split(string(input), "\n")
  slope := 3
  result := DayThreePartOne(lines, slope)
  fmt.Println("3.1:", result)
  slopes := []float64{1}
  result = DayThreePartTwo(lines, slopes)
  fmt.Printf("3.2: %v (Incorrect)\n", result)
}

func DayThreePartOne(data []string, slope int) (tree_count int) {
  width := len(data[0])
  pos_x := 0
  for _, row := range data { 
    if row[pos_x] == '#' {
      tree_count += 1
    }
    pos_x += slope
    pos_x %= width
  }
  return tree_count
}

func DayThreePartTwo(data []string, slopes []float64) (tree_product int) {
  // TODO: I'm too tired. 
  // 1 - 84
  // 3 - 195
  // 5 - 70
  // 7 - 70
  // 0.5 - 47 This is manually calculated by the below. feelsbad
  width := float64(len(data[0]))
  pos_x := float64(0)
  tree_product = 1
  for _, slope := range slopes {
    tree_count := 0
    for cop_out, row := range data {
      if cop_out % 2 != 0  { 
        continue
      }
      floor_pos := math.Floor(pos_x)
      if (math.Ceil(pos_x) == floor_pos) && (row[int(floor_pos)] == '#') {
        tree_count += 1
        // fmt.Print("X ")
      } else {
        // fmt.Print("  ")
      }
      // fmt.Println(row, pos_x)
      pos_x += slope
      if math.Floor(pos_x) >= width {
        pos_x -= width
      }
    }
    tree_product *= tree_count
  }
  return tree_product
}

func AoCDayTwo() {
  input, _ := ioutil.ReadFile("day_2_input.txt")
  lines := strings.Split(string(input), "\n")
  result := DayTwoPartOne(lines)
  fmt.Println("2.1:", result)
  result = DayTwoPartTwo(lines)
  fmt.Println("2.2:", result)
}

func DayTwoPartOne(data []string) (valid_count int) {
  for _, line := range data {
    // Sample Input: "9-15 d: dqddhddngdddddnzd"
    result := strings.Split(line, " ")
    freq_arr := strings.Split(result[0], "-") // Gets ["9", "15"] from "9-15"

    min_freq, err := strconv.Atoi(freq_arr[0]) // Gets "9"
    if err != nil {
      valid_count = -1
      return 
    }
    max_freq, err := strconv.Atoi(freq_arr[1]) // Gets "15"
    if err != nil {
      valid_count = -1
      return 
    }
    char := result[1][0:] // Gets "d" from "d:"
    s := result[2] // Gets "dqddhddngdddddnzd"

    chars_seen := 0
    for _, c := range s {
      s, _ := utf8.DecodeRuneInString(char)
      if c == s {
        chars_seen += 1
      }
    }
    
    if min_freq <= chars_seen && chars_seen <= max_freq {
      valid_count += 1
    }
  }
  return
}

func DayTwoPartTwo(data []string) (valid_count int) {
  for _, line := range data {
    result := strings.Split(line, " ")
    pos_arr := strings.Split(result[0], "-") 
    pos_a, err := strconv.Atoi(pos_arr[0])
    if err != nil {
      valid_count = -1
      return
    }
    pos_b, err := strconv.Atoi(pos_arr[1])
    if err != nil {
      valid_count = -1
      return
    }
    char := result[1][0]
    s := result[2]

    ix_a := pos_a - 1
    ix_b := pos_b - 1

    if (s[ix_a] == char || s[ix_b] == char) && s[ix_a] != s[ix_b] {
      valid_count += 1
    }
  }
  return 
}

func AoCDayOne() {
  input, _ := ioutil.ReadFile("day_1_input.txt")
  lines := strings.Split(string(input), "\n")
  target := 2020
  result := DayOnePartOne(lines, target)
  fmt.Println("1.1:", result)
  result = DayOnePartTwo(lines, target)
  fmt.Println("1.2", result)
}

func DayOnePartOne(data []string, target int) (product int) {
  product = -1
  find_vals := make(map[int]struct{})
  for _, line := range data {
    num, err := strconv.Atoi(line)
    if err != nil {
      break
    }
    if _, ok := find_vals[num]; ok {
      product = num * (target-num)
      break
    }
    find_vals[target-num] = struct{}{}
  }
  return 
}

func DayOnePartTwo(data []string, target int) (product int) {
  product = -1
  find_vals := make(map[int]int)
  for ix, val_a := range data {
    for _, val_b := range data[ix:] {
      num_a, err := strconv.Atoi(val_a)
      if err != nil {
        return
      }
      num_b, err := strconv.Atoi(val_b)
      if err != nil {
        return
      }
      if _, ok := find_vals[num_b]; ok {
        product = find_vals[num_b] * num_b
        return
      }
      find_vals[target-num_a-num_b] = num_a * num_b
    }
  }
  return
}
