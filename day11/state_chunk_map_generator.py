# def generate_states(s='', iter=1):
#   if remaining == 0:
#     return ['']
#   results = []
#   for c in ['.', 'L', '#']:
#     for n in generate_states(c, iter - 1):
#       results.append(c+n)
#   return results
#
# def get_derived_chunk(state):
#     return ''.join(get_next_state(state, pos) for pos in [5, 6, 9, 10])
#
# def get_next_state(state, pos):
#     star = state[pos]
#     if star == '.':
#         return '.'
#     elif star == 'L':
#         for adj_pos in [
#             pos-5, pos-4, pos-3,
#             pos-1,        pos+1,
#             pos+3, pos+4, pos+5,
#         ]:
#             if state[adj_pos] == '#': # found occupied seat
#                 return 'L'
#         return '#' # no occupied seats around empty seat, so it's occupied now
#     elif star == '#':
#         filled_adj_seats = 0
#         for adj_pos in [
#             pos-5, pos-4, pos-3,
#             pos-1,        pos+1,
#             pos+3, pos+4, pos+5,
#         ]:
#             if state[adj_pos] == '#':
#                 filled_adj_seats += 1
#             if filled_adj_seats == 4:
#                 return 'L' # seat is deserted, found 4 adjacent occupied
#         return '#' # seat is still occupied, didn't find 4 adjacent occupied
#
# with open('/Users/alo/go/src/aoc/state_chunk_map_2.csv', 'a') as fi:
#     for state in recursion('', 16):
#         fi.write(f"{state},{get_derived_chunk(state)}\n")
#
# get_next_state = (
#     lambda state, pos: \
#         '.' if state[pos] == '.' else \
#
#         (
#             'L' if 0 < len([ # if len > 0, found an occupied adjacent seat, so this seat is still unoccupied
#                 1
#                 for adj_pos in [
#                     pos-5, pos-4, pos-3,
#                     pos-1,        pos+1,
#                     pos+3, pos+4, pos+5,
#                 ] if state[adj_pos] == '#'
#             ]) else \
#             '#' # seat is now occupied, didn't find any adjacent seats
#         ) if state[pos] == 'L' else \
#
#         (
#             'L' if 4 <= len([ # if len >= 4, found 4 adjacent occupied, so this seat becomes unoccupied
#                 1
#                 for adj_pos in [
#                     pos-5, pos-4, pos-3,
#                     pos-1,        pos+1,
#                     pos+3, pos+4, pos+5,
#                 ] if state[adj_pos] == '#'
#             ]) else
#         ) if state[pos] == '#' # didn't find enough adjacent occupied, so seat is still occupied
# )
#
# generate_states = lambda iter: \
#     [''] if iter == 0 else \
#     [
#         state
#         for list_of_states in [
#             [
#                 c + n
#                 for n in generate_states(iter - 1)
#             ] for c in ('.', 'L', '#')
#         ] for state in list_of_states
#     ]
#
# # This is so cool
# # Source: https://en.wikipedia.org/wiki/Anonymous_recursion
# F = lambda f: (lambda x: f(f, x))
# gens1 = lambda f, iter: [''] if iter == 0 else \
#     [
#         state
#         for list_of_states in [
#             [
#                 c + n
#                 for n in f(f, iter - 1)
#             ] for c in ('.', 'L', '#')
#         ] for state in list_of_states
#     ]
result = (lambda f: (lambda x: f(f, x)))(lambda gen_states, iter: [''] if iter == 0 else [state for list_of_states in [[char + next_char for next_char in gen_states(gen_states, iter - 1)] for char in '.L#'] for state in list_of_states])
print(result(4))
#
# open('./state_chunk_map_2.csv', 'w').write(
print(
    "\n".join([
        "{state},{derived_state}".format(
            state=state,
            derived_state=''.join(
                [
                    (
                        lambda state, pos: \
                            "." if state[pos] == "." else \

                            (
                                "L" if 0 < len([ # if len > 0, found an occupied adjacent seat, so this seat is still unoccupied
                                    1
                                    for adj_pos in [
                                        pos-5, pos-4, pos-3,
                                        pos-1,        pos+1,
                                        pos+3, pos+4, pos+5,
                                    ] if state[adj_pos] == "#"
                                ]) else \
                                "#" # seat is now occupied, didn't find any adjacent seats
                            ) if state[pos] == "L" else \
                            (
                                "L" if 4 <= len([ # if len >= 4, found 4 adjacent occupied, so this seat becomes unoccupied
                                    1
                                    for adj_pos in [
                                        pos-5, pos-4, pos-3,
                                        pos-1,        pos+1,
                                        pos+3, pos+4, pos+5,
                                    ] if state[adj_pos] == "#"
                                ]) else "#"
                            ) # if state[pos] == '#' # didn't find enough adjacent occupied, so seat is still occupied
                    )(state, pos)
                    for pos in [5, 6, 9, 10] # Center squares
                ]
            )
        ) for state in result(16)
    ])
)
#
# # Could possibly reduce generate_states into manual 16-loop nested list comprehension, but maybe another day
# # generate_states = lambda iter: [''] if iter == 0 else [state for list_of_states in [[c + n for n in generate_states(iter - 1)] for c in ('.', 'L', '#')] for state in list_of_states]
open("./state_chunk_map_2.csv", "w").write("\n".join(["{state},{derived_state}".format(state=state, derived_state="".join([(lambda state, pos: "." if state[pos] == "." else ("L" if 0 < len([1 for adj_pos in [pos-5, pos-4, pos-3, pos-1, pos+1, pos+3, pos+4, pos+5] if state[adj_pos] == "#"]) else "#") if state[pos] == "L" else ("L" if 4 <= len([1 for adj_pos in [pos-5, pos-4, pos-3, pos-1, pos+1, pos+3, pos+4, pos+5] if state[adj_pos] == "#"]) else "#"))(state, pos) for pos in [5, 6, 9, 10]])) for state in (lambda f: (lambda x: f(f, x)))(lambda gen_states, iter: [""] if iter == 0 else [state for list_of_states in [[char + next_char for next_char in gen_states(gen_states, iter - 1)] for char in ".L#"] for state in list_of_states])(2)]))

