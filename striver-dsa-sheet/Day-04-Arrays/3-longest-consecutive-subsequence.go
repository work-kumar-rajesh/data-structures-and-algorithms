func longestConsecutive(nums []int) int {
    numsSet := make(map[int]struct{})
    for _, n := range nums {
        numsSet[n] = struct{}{}
    }

    maxSequenceLen := 0

    for n := range numsSet {
        if _, ok := numsSet[n-1]; !ok {
            seqLen := 1
            for {
                if _, ok = numsSet[n+seqLen]; ok {
                    seqLen++
                    continue
                }
                maxSequenceLen = max(seqLen, maxSequenceLen)
                break
            }
        }
    }

    return maxSequenceLen
}
