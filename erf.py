import math

#'Cumulative distribution function for the standard normal distribution'

def phi(x):

    return (1.0 + erf(x / sqrt(2.0))) / 2.0
