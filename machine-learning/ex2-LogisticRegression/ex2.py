# coding=utf-8
import matplotlib.pyplot as plt
import numpy as np
import scipy.optimize as opt


def plotData(X, y):
    """
    Plots the data points X and y into a new figure
    :param X:
    :param y:
    :return:
    """
    pos = np.argwhere(y == 1)
    neg = np.argwhere(y == 0)
    plt.plot(X[pos, 0], X[pos, 1], linestyle="", marker="+", color="k")
    plt.plot(X[neg, 0], X[neg, 1], linestyle="", marker="o", color="y")


def sigmoid(z):
    """
    Compute sigmoid function
    :param z:
    :return:
    """
    g = 1 / (1 + np.exp(-z))
    return g


def costFunction(theta, X, y):
    """
    Compute cost and gradient for logistic regression
    :param theta:
    :param X:
    :param y:
    :return:
    """
    m = X.shape[0]
    J = (1.0 / m) * np.sum((-y * np.log(sigmoid(np.dot(X, theta)))) - ((1 - y) * np.log(1 - sigmoid(np.dot(X, theta)))))
    grad = (1.0 / m) * (np.dot(np.transpose(X), (sigmoid(np.dot(X, theta)) - y)))
    return J, grad


def plotDecisionBoundary(theta, X, y):
    """
    Plots the data points X and y into a new figure with the decision boundary defined by theta
    :param theta:
    :param X:
    :param y:
    :return:
    """
    plotData(X[:, 1:3], y)
    if X.shape[1] <= 3:
        plot_X = np.array([np.amin(X[:, 1]) - 2, np.amax(X[:, 1]) + 2])
        plot_y = -1.0 / theta[2] * (theta[1] * plot_X + theta[0])
        plt.plot(plot_X, plot_y)
        plt.legend(['Decision Boundary', 'Admitted', 'Not admitted'], loc="upper right")
        plt.axis([30, 100, 30, 100])
    else:
        u = np.linspace(-1, 1.5, 50)
        v = np.linspace(-1, 1.5, 50)
        z = np.zeros((len(u), len(v)))
        for i in range(len(u)):
            for j in range(len(v)):
                z[i, j] = mapFeature(u[i:i + 1], v[j:j + 1]).dot(theta)
        z = z.T
        u, v = np.meshgrid(u, v)
        cs = plt.contour(u, v, z, levels=[0])
        fmt = {}
        strs = ['Decision boundary']
        for l, s in zip(cs.levels, strs):
            fmt[l] = s
        plt.clabel(cs, cs.levels[::2], inline=True, fmt=fmt, fontsize=10)


def mapFeature(X1, X2):
    """
    Feature mapping function to polynomial features
    :param X1:
    :param X2:
    :return:
    """
    degree = 6
    Out = np.ones(len(X1))
    for i in range(1, degree + 1):
        for j in range(0, i + 1):
            tmp = np.power(X1, i - j) * np.power(X2, j)
            Out = np.vstack((Out, tmp))
    return Out.T


def predict(theta, X):
    """
    Predict whether the label is 0 or 1 using learned logistic regression parameters theta
    :param theta:
    :param X:
    :return:
    """
    p = sigmoid(X.dot(theta)) >= 0.5
    return p.astype(int)


def costFunctionReg(theta, X, y, labda):
    """
    Compute cost and gradient for logistic regression with regularization
    :param theta:
    :param X:
    :param y:
    :param labda:
    :return:
    """
    m = X.shape[0]
    e = np.eye(len(theta))
    e[0, 0] = 0
    J = (1.0 / m) * np.sum((-y * np.log(sigmoid(np.dot(X, theta)))) - ((1 - y) * np.log(1 - sigmoid(np.dot(X, theta))))) + (labda / 2 * m) * np.sum(
        np.power(theta.dot(e), 2))
    grad = (1.0 / m) * (np.dot(np.transpose(X), (sigmoid(np.dot(X, theta)) - y))) + (labda / m) * theta.dot(e)
    return J, grad


if __name__ == "__main__":
    # load data
    data = np.loadtxt(r'ex2data1.txt', delimiter=',')
    xData = data[:, 0:2]
    yData = data[:, 2]
    m = data.shape[0]

    # ==================== Part 1: Plotting ====================
    print('Plotting data with + indicating (y = 1) examples and o indicating (y = 0) examples.')
    plotData(xData, yData)
    plt.xlabel('Exam 1 score')
    plt.ylabel('Exam 2 score')
    plt.legend(['Admitted', 'Not admitted'], loc="upper right")
    plt.show()

    # ============ Part 2: Compute Cost and Gradient ============
    X = np.c_[np.ones(m), xData]
    y = yData
    theta = np.zeros(X.shape[1])
    cost, grad = costFunction(theta, X, y)
    print('Cost at initial theta (zeros): ', cost)
    print('Expected cost (approx): 0.693')
    print('Gradient at initial theta (zeros): \n', grad)
    print('Expected gradients (approx): \n[-0.1000 -12.0092 -11.2628]')

    # Compute and display cost and gradient with non-zero theta
    test_theta = np.array([-24, 0.2, 0.2])
    cost, grad = costFunction(test_theta, X, y)

    print('Cost at test theta: ', cost)
    print('Expected cost (approx): 0.218')
    print('Gradient at test theta: \n', grad)
    print('Expected gradients (approx):\n [0.043 2.566 2.647]')

    # ============= Part 3: Optimizing using fminunc  =============
    theta, nfeval, rc = opt.fmin_tnc(func=costFunction, x0=theta, args=(X, y), messages=0)
    if rc == 0:
        print('Local minimum reached after {} function evaluations.'.format(nfeval))
    cost, _ = costFunction(theta, X, y)
    print('Cost at theta found by fminunc: ', cost)
    print('Expected cost (approx): 0.203')
    print('theta: \n', theta)
    print('Expected theta (approx): \n [-25.161 0.206 0.201]')

    # Plot Boundary
    plotDecisionBoundary(theta, X, y)
    plt.show()

    # ============== Part 4: Predict and Accuracies ==============
    prob = sigmoid(np.array([1, 45, 85]).dot(theta))
    print('For a student with scores 45 and 85, we predict an admission probability of ', prob)
    print('Expected value: 0.775 +/- 0.002')

    # Compute accuracy on our training set
    p = predict(theta, X)
    print('Train Accuracy: ', np.mean(p == y) * 100)
    print('Expected accuracy (approx): 89.0\n\n')

    # ============================= Regularized logistic regression ===============================
    # Load Data
    data = np.loadtxt(r'ex2data2.txt', delimiter=",")
    X = data[:, 0:2]
    y = data[:, 2]
    m = data.shape[0]

    plotData(X, y)
    plt.xlabel('Microchip Test 1')
    plt.ylabel('Microchip Test 2')
    plt.legend(['y = 1', 'y = 0'], loc="upper right")
    plt.show()

    X = mapFeature(X[:, 0], X[:, 1])
    initial_theta = np.zeros(X.shape[1])
    labda = 1
    cost, grad = costFunctionReg(initial_theta, X, y, labda)
    print('Cost at initial theta (zeros): ', cost)
    print('Expected cost (approx): 0.693')
    print('Gradient at initial theta (zeros) - first five values only: \n', grad[0:5])
    print('Expected gradients (approx) - first five values only:')
    print('[0.0085 0.0188 0.0001 0.0503 0.0115]\n')
    test_theta = np.ones(X.shape[1])
    labda = 10
    cost, grad = costFunctionReg(test_theta, X, y, labda)
    print('Cost at test theta (with lambda = 10): ', cost)
    print('Expected cost (approx): 3.16')
    print('Gradient at test theta - first five values only:\n', grad[0:5])
    print('Expected gradients (approx) - first five values only:')
    print('[0.3460 0.1614 0.1948 0.2269 0.0922]\n')

    initial_theta = np.zeros(X.shape[1])
    labda = 1
    theta, nfeval, rc = opt.fmin_tnc(func=costFunctionReg, x0=initial_theta, args=(X, y, labda), messages=0)
    if rc == 0:
        print('Local minimum reached after {} function evaluations.'.format(nfeval))
    plotDecisionBoundary(theta, X, y)
    plt.title("lambda = {}".format(labda))
    plt.xlabel('Microchip Test 1')
    plt.ylabel('Microchip Test 2')
    plt.legend(['y = 1', 'y = 0'], loc="upper right")
    plt.show()

    # Compute accuracy on our training set
    p = predict(theta, X)
    print('Train Accuracy: ', np.mean(p == y) * 100)
    print('Expected accuracy (with lambda = 1): 83.1 (approx)')
