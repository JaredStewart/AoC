def is_unsafe(vals):
    is_increase = False
    last_val = None
    delta = None
    unsafe = False
    for val in vals:
        if last_val is None:
            last_val = val
            continue
        if delta is None:
            is_increase = (val - last_val) > 0
        
        delta = val- last_val
        last_val = val
        if delta == 0:
            unsafe = True

        if delta < 0:
            if is_increase:
                unsafe = True
            if delta < -3:
                unsafe = True

        if delta > 0:
            if not is_increase:
                unsafe = True
            if delta > 3:
                unsafe = True

    return unsafe

with open('input.txt') as f:
    pure_safe_count = 0
    backup_safe_count = 0
    for line in f:
        vals = [int(v) for v in line.split()]
        if is_unsafe(vals):
            for vals_2 in [vals[:i] + vals[i+1:] for i in range(len(vals))]:
                if not is_unsafe(vals_2):
                    print(vals)
                    backup_safe_count += 1
                    break

        else:
            pure_safe_count += 1
        

    print(pure_safe_count)
    print(backup_safe_count)
