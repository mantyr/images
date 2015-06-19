package image

import (
    "math"
)

type ResampleFilter struct {
    Support float64
    Kernel  func(float64) float64
}

var NearestNeighbor ResampleFilter
var Box ResampleFilter
var Linear ResampleFilter
var Hermite ResampleFilter
var MitchellNetravali ResampleFilter
var CatmullRom ResampleFilter
var BSpline ResampleFilter
var Gaussian ResampleFilter
var Bartlett ResampleFilter
var Lanczos ResampleFilter
var Hann ResampleFilter
var Hamming ResampleFilter
var Blackman ResampleFilter
var Welch ResampleFilter
var Cosine ResampleFilter

func bcspline(x, b, c float64) float64 {
    x = math.Abs(x)
    if x < 1.0 {
        return ((12-9*b-6*c)*x*x*x + (-18+12*b+6*c)*x*x + (6 - 2*b)) / 6
    }
    if x < 2.0 {
        return ((-b-6*c)*x*x*x + (6*b+30*c)*x*x + (-12*b-48*c)*x + (8*b + 24*c)) / 6
    }
    return 0
}

func sinc(x float64) float64 {
    if x == 0 {
        return 1
    }
    return math.Sin(math.Pi*x) / (math.Pi * x)
}

func init() {
    NearestNeighbor = ResampleFilter{
        Support: 0.0, // special case - not applying the filter
    }

    Box = ResampleFilter{
        Support: 0.5,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x <= 0.5 {
                return 1.0
            }
            return 0
        },
    }

    Linear = ResampleFilter{
        Support: 1.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 1.0 {
                return 1.0 - x
            }
            return 0
        },
    }

    Hermite = ResampleFilter{
        Support: 1.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 1.0 {
                return bcspline(x, 0.0, 0.0)
            }
            return 0
        },
    }

    MitchellNetravali = ResampleFilter{
        Support: 2.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 2.0 {
                return bcspline(x, 1.0/3.0, 1.0/3.0)
            }
            return 0
        },
    }

    CatmullRom = ResampleFilter{
        Support: 2.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 2.0 {
                return bcspline(x, 0.0, 0.5)
            }
            return 0
        },
    }

    BSpline = ResampleFilter{
        Support: 2.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 2.0 {
                return bcspline(x, 1.0, 0.0)
            }
            return 0
        },
    }

    Gaussian = ResampleFilter{
        Support: 2.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 2.0 {
                return math.Exp(-2 * x * x)
            }
            return 0
        },
    }

    Bartlett = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * (3.0 - x) / 3.0
            }
            return 0
        },
    }

    Lanczos = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * sinc(x/3.0)
            }
            return 0
        },
    }

    Hann = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * (0.5 + 0.5*math.Cos(math.Pi*x/3.0))
            }
            return 0
        },
    }

    Hamming = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * (0.54 + 0.46*math.Cos(math.Pi*x/3.0))
            }
            return 0
        },
    }

    Blackman = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * (0.42 - 0.5*math.Cos(math.Pi*x/3.0+math.Pi) + 0.08*math.Cos(2.0*math.Pi*x/3.0))
            }
            return 0
        },
    }

    Welch = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * (1.0 - (x * x / 9.0))
            }
            return 0
        },
    }

    Cosine = ResampleFilter{
        Support: 3.0,
        Kernel: func(x float64) float64 {
            x = math.Abs(x)
            if x < 3.0 {
                return sinc(x) * math.Cos((math.Pi/2.0)*(x/3.0))
            }
            return 0
        },
    }
}
