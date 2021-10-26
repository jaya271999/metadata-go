func calcSingleCoreUsage(curr, prev linuxproc.CPUStat) float32 {

  PrevIdle := prev.Idle + prev.IOWait
  Idle := curr.Idle + curr.IOWait

  PrevNonIdle := prev.User + prev.Nice + prev.System + prev.IRQ + prev.SoftIRQ + prev.Steal
  NonIdle := curr.User + curr.Nice + curr.System + curr.IRQ + curr.SoftIRQ + curr.Steal

  PrevTotal := PrevIdle + PrevNonIdle
  Total := Idle + NonIdle
  // fmt.Println(PrevIdle, Idle, PrevNonIdle, NonIdle, PrevTotal, Total)

  //  differentiate: actual value minus the previous one
  totald := Total - PrevTotal
  idled := Idle - PrevIdle

  CPU_Percentage := (float32(totald) - float32(idled)) / float32(totald)

  return CPU_Percentage
}
