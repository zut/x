package xx

import (
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/lapack/lapack64"
	"gonum.org/v1/gonum/mat"
	"math"
)

func WhereGtSI(s []int, limit, a, b int) []int {
	d := make([]int, len(s))
	for n, i := range s {
		d[n] = IfInt(i > limit, a, b)
	}
	return d
}
func WhereGtNum(s []float64, limit float64) (num int) {
	for n := range s {
		if s[n] > limit {
			num++
		}
	}
	return
}
func WhereGT(s []float64, limit, a, b float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i > limit, a, b)
	}
	return d
}
func WhereGT2(s []float64, limit float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i > limit, i, 0)
	}
	return d
}
func WhereGE(s []float64, limit, a, b float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i >= limit, a, b)
	}
	return d
}

func WhereLT(s []float64, limit, a, b float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i < limit, a, b)
	}
	return d
}

func WhereLtNum(s []float64, limit float64) (num int) {
	for n := range s {
		if s[n] < limit {
			num++
		}
	}
	return
}

// SumWhereGT 2维变1维，然后相加
func SumWhereGT(src [][]float64, limit, a, b float64) []float64 {
	if len(src) == 0 {
		return nil
	}
	dst := make([]float64, len(src[0]))
	valueListOfPerIdx := make([]float64, len(src))
	for n := range src[0] {
		for idx := range valueListOfPerIdx {
			valueListOfPerIdx[idx] = src[idx][n]
		}
		dst[n] = IfF64(floats.Sum(valueListOfPerIdx) > limit, a, b)
	}
	return dst
}

// Split01 一维变二维
func Split01(src []float64) [][]float64 {
	if len(src) == 0 {
		return nil
	}
	dst := make([][]float64, 0)
	last := -1.0
	lastIdx := 0
	endN := len(src) - 1
	for n, i := range src {
		//if n == len(src)-2 {
		//	xlog.Info(i)
		//}
		if last != i { // 不同,改变了
			if n != 0 { // 第一个不要进来
				dst = append(dst, src[lastIdx:n])
			}
			last = i
			lastIdx = n
			if n == endN { // 最后一个只有1个的时候, 也需要加入, 否者漏点1个
				dst = append(dst, src[lastIdx:])
			}
		} else {
			if n == endN {
				dst = append(dst, src[lastIdx:])
			}
		}
	}
	return dst
}

// PullUp01 拉点
func PullUp01(src [][]float64, point int) [][]float64 {
	if len(src) == 0 {
		return nil
	}
	dst := make([][]float64, 0)
	for n, i := range src {
		if i[0] == 0 && len(i) < point {
			for m := range i {
				i[m] = 1
			}
		}
		if n == 0 || dst[len(dst)-1][0] != i[0] {
			dst = append(dst, i)
		} else {
			dst[len(dst)-1] = append(dst[len(dst)-1], i...)
		}
	}
	return dst
}

// other

// Zeroes returns an array of zeroes of specified size.
// It's encouraged to use it instead of just make() in case the further code relies on the fact that the array contains zeroes.
func Zeroes(size int) []float64 {
	return make([]float64, size)
}

// Ones return an array of ones of specified size.
func Ones(size int) []float64 {
	result := make([]float64, size)
	for i := range result {
		result[i] = 1
	}
	return result
}

// LinSpace implements `np.LinSpace` - i.e. splits the interval [start, end] into `num - 1` equal intervals and returns `num` split points.
func LinSpace(start, end float64, num int) []float64 {
	if num < 0 {
		panic(errors.Errorf("number of samples, %d, must be non-negative.", num))
	}
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result
}

// ARange implements `np.ARange` - i.e. returns a list of integers (start, ..., stop - 1) in the form of []float64
func ARange(start int, stop int) []float64 {
	return LinSpace(float64(start), float64(stop-1), stop-start)
}

// ConvolutionMode defines convolution output array length.
type ConvolutionMode int

const (
	// Full - returns the convolution at each point of overlap, i.e. of length N+M-1.
	Full = iota
	// Same - returns the output of length max(M, N).
	Same
	// Valid - returns the output of length max(M, N) - min(M, N) + 1.
	Valid
)

// Convolve is a (very naive) implementation of precise discrete convolution.
// The results are numerically equivalent to `np.convolve(a, v, mode)` - it looks like that `np.convolve` uses precise convolution as well (but not an FFT approximation).
// TODO: optimize the implementation - the current one has O((M+N)^2) time complexity. Looks like it's possible to achieve at least O(MN).
func Convolve(a, v []float64, mode ConvolutionMode) []float64 {
	if len(a) == 0 {
		panic(errors.New("Convolve: a cannot be empty"))
	}
	if len(v) == 0 {
		panic(errors.New("Convolve: v cannot be empty"))
	}

	// the code below relies on the fact that `a` is the longer array
	if len(v) > len(a) {
		a, v = v, a
	}

	size := len(a) + len(v) - 1

	// a + zeroes
	aExt := Zeroes(size)
	copy(aExt[:len(a)], a)

	// v + zeroes
	vExt := Zeroes(size)
	copy(vExt[:len(v)], v)

	result := Zeroes(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			result[i+j] += aExt[i] * vExt[j]
		}
	}

	switch mode {
	case Full:
		// `Full` mode: returning the whole result
		return result
	case Same:
		// `Same` mode: returning the subArray of length `len(a)`
		toCut := len(v) / 2 // is this correct? at least, this works for the sample data
		return result[toCut : toCut+len(a)]
	case Valid:
		// `Valid` mode: returning the subArray of length `len(a) - len(v) + 1`
		toCut := len(v) - 1
		return result[toCut:len(a)]
	default:
		panic(errors.Errorf("invalid convolution mode %v", mode))
	}
}

// computes `n!`
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// SavGolFilter implements Savitzky-Golay filter (https://docs.scipy.org/doc/scipy/reference/generated/scipy.signal.savgol_filter.html)
// based on: https://github.com/scipy/scipy/blob/v1.3.0rc1/scipy/signal/_savitzky_golay.py#L227
func SavGolFilter(x []float64, windowLength int, polyOrder int, deRiv int /*=0*/, delta float64 /*=1.0*/) []float64 {
	windowLength = MaxInt(SI{windowLength, 2})
	if windowLength%2 == 0 {
		windowLength = windowLength + 1
	}
	// computing filter coefficients
	// the outputs of this step seem to be numerically same as the Python code ones
	coEffs := savGolCoEffs(windowLength, polyOrder, deRiv, delta, true)
	// convOlVing the original signal with the filter coefficients
	// note: the outputs of this step are not completely numerically same as the Python code ones (because the latter uses different convolution function)
	return Convolve(x, coEffs, Same)
}

// Computes Savitzky-Golay filter coefficients.
// based on: https://github.com/scipy/scipy/blob/v1.3.0rc1/scipy/signal/_savitzky_golay.py#L10
func savGolCoEffs(windowLength int, polyOrder int, deRiv int, delta float64, useInConv bool) []float64 {
	if polyOrder >= windowLength {
		panic("polyorder must be less than window_length.")
	}
	if windowLength%2 == 0 {
		panic("window_length must be odd.")
	}
	pos := windowLength / 2
	if pos < 0 {
		panic("pos must be nonnegative.")
	}

	// Form the design matrix `A`. The columns of `A` are powers of the integers
	// from -pos to window_length - pos - 1.  The powers (i.e. rows) range
	// from 0 to polyOrder.
	aRowTemplate := ARange(-pos, windowLength-pos)
	if useInConv {
		// Reverse so that result can be used in a convolution.
		floats.Reverse(aRowTemplate)
	}
	a := makeMatrix(polyOrder+1, len(aRowTemplate), func(i, j int) float64 {
		return math.Pow(aRowTemplate[j], float64(i))
	})

	// `b` determines which order derivative is returned.
	// The coefficient assigned to b[deRiv] scales the result to take into
	// account the order of the derivative and the sample spacing.
	b := makeMatrix(polyOrder+1, 1, func(i, j int) float64 {
		if i != deRiv {
			return 0
		}
		return float64(factorial(deRiv)) / math.Pow(delta, float64(deRiv))
	})

	// finding the least-squares solution of A*x = b
	coEff := LstSq(a, b)
	if _, cols := coEff.Dims(); cols != 1 {
		panic(errors.Errorf("SHOULD NOT HAPPEN: LstSq result contains %d columns instead of 1", cols))
	}
	return coEff.RawMatrix().Data
}

// Makes a dense matrix of size r*c and fills it with a user-defined function.
func makeMatrix(r int, c int, value func(i, j int) float64) *mat.Dense {
	data := make([]float64, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			data[c*i+j] = value(i, j)
		}
	}
	return mat.NewDense(r, c, data)
}

// LstSq computes least-squares solution to equation A*x = b, i.e. computes a vector x such that the 2-norm “|b - A x|“ is minimized.
// Based on: https://github.com/scipy/scipy/blob/v1.3.0rc1/scipy/linalg/basic.py#L1042
func LstSq(a, b *mat.Dense) *mat.Dense {
	// m is a number of columns in `a`, n is a number of rows in `a`
	m, n := a.Dims()
	if m == 0 || n == 0 {
		panic("zero-sized problem is not supported (confuses LAPACK)")
	}

	m2, nhRs := b.Dims()
	if m2 != m {
		panic(errors.Errorf("shape mismatch: a and b should have the same number of rows: %d != %d", m, m2))
	}

	// LAPACK uses `b` as an output parameter as well - and therefore wants it to be resized from (m, nhrs) to (max(m,n), nhrs)
	// here we copy `b` anyway (even if it doesn't need to be resized) - to avoid overwriting the user-supplied `b`
	b = makeMatrix(MaxInt(SI{m, n}), nhRs, func(i, j int) float64 {
		if i < m {
			return b.At(i, j)
		}
		return 0
	})

	// LAPACK function for computing least-squares solutions to linear equations
	gels := func(work []float64, lwork int) bool {
		return lapack64.Gels(blas.NoTrans, a.RawMatrix(), b.RawMatrix(), work, lwork)
	}

	// retrieving the size of work space needed (this is how LAPACK interfaces are designed:
	// if we call the function with lwork=-1, it returns the work size needed in work[0])
	work := make([]float64, 1)
	gels(work, -1)
	lWork := int(math.Ceil(work[0]))

	// solving the equation itself
	result := gels(make([]float64, lWork), lWork)
	if !result {
		panic(errors.Errorf("gels: computation didn't converge: A='%+v', b='%+v'", a, b))
	}

	return b
}
