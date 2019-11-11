# coding=utf-8
import matplotlib
import matplotlib.pyplot as plt
import numpy as np
from matplotlib.colors import LogNorm
from mpl_toolkits.mplot3d import Axes3D

matplotlib.get_cachedir()  # fix `Unable to revert mtime: /Library/Fonts`


def warmUpExercise():
    """
    :return: 5x5 单位矩阵(identity matrix)
    """
    return np.eye(5)


def plotData(x, y):
    """
    绘制训练数据散点图
    :param x:
    :param y:
    :return:
    """
    plt.scatter(x, y, s=50, c="red", marker="x", alpha=0.6)
    plt.title("Scatter plot")
    plt.xlabel("Population of City in 10,000s")
    plt.ylabel("Profit in $10,000s")
    plt.show()


def computeCost(X, y, theta):
    """
    Compute cost for linear regression
    :param X:
    :param y:
    :param theta:
    :return:
    """
    cost = np.dot(X, theta) - y
    return np.sum(np.power(cost, 2)) / (2 * y.shape[0])


def gradientDescent(X, y, theta, alpha, iterations):
    """
    Performs gradient descent to learn theta
    :param X:
    :param y:
    :param theta:
    :param alpha:
    :param iterations:
    :return:
    """
    m = y.shape[0]
    J_history = np.zeros((iterations, 1))
    O_theta = theta
    for i in range(iterations):
        O_theta = O_theta - (alpha / m) * (np.dot(np.transpose(X), (np.dot(X, O_theta) - y)))
        J_history[i] = computeCost(X, y, O_theta)
    return O_theta, J_history


if __name__ == "__main__":
    # ==================== Part 1: Basic Function ====================
    print("Running warmUpExercise ... ")
    I = warmUpExercise()
    print("5x5 Identity Matrix I{}: \n{}".format(type(I), I))

    # ======================= Part 2: Plotting =======================
    print("Plotting Data ...")
    data = np.loadtxt(r"ex1data1.txt", delimiter=",", usecols=(0, 1))  # numpy.ndarray
    xData = data[:, 0]
    yData = data[:, 1]
    m = data.shape[0]
    print("number of training examples: ", m)
    plotData(xData, yData)

    # =================== Part 3: Cost and Gradient descent ===================
    X = np.c_[np.ones(m), data[:, 0]]  # Add a column of ones to x; now X.shape is (m, 2)
    Y = yData.reshape((m, 1))
    theta = np.zeros((2, 1))  # initialize fitting parameters
    # Some gradient descent settings
    iterations = 1500
    alpha = 0.01

    print("Testing the cost function ...")
    # compute and display initial cost
    J = computeCost(X, Y, theta)
    print('With theta = [0 ; 0]\nCost computed = {}'.format(J))
    print('Expected cost value (approx) 32.07')

    # further testing of the cost function
    J = computeCost(X, Y, np.array([[-1], [2]]))
    print('With theta = [-1 ; 2]\nCost computed = {}'.format(J))
    print('Expected cost value (approx) 54.24')

    print('Running Gradient Descent ...')
    # run gradient descent
    theta, J_history = gradientDescent(X, Y, theta, alpha, iterations)
    # print theta to screen
    print('Theta found by gradient descent:')
    print(theta)
    print('Expected theta values (approx)')
    print(' -3.6303\n  1.1664\n')

    # Plot the linear fit
    plt.figure(0)
    plt.xlabel("Population of City in 10,000s")
    plt.ylabel("Profit in $10,000s")
    plt.scatter(xData, yData, c="red", marker="x", label="Training data")
    plt.plot(X[:, 1], np.dot(X, theta), label="Linear regression")
    plt.legend(loc="best")
    plt.show()

    # Predict values for population sizes of 35,000 and 70,000
    predict1 = np.dot(np.array([[1, 3.5]]), theta)
    print('For population = 35,000, we predict a profit of ', predict1 * 10000)
    predict2 = np.dot(np.array([[1, 7]]), theta)
    print('For population = 70,000, we predict a profit of ', predict2 * 10000)

    # ============= Part 4: Visualizing J(theta_0, theta_1) =============
    # Grid over which we will calculate J
    theta0_vals = np.linspace(-10, 10, 100)
    theta1_vals = np.linspace(-1, 4, 100)
    # initialize J_vals to a matrix of 0's
    J_vals = np.zeros((theta0_vals.shape[0], theta1_vals.shape[0]))
    # Fill out J_vals
    for i in range(0, theta0_vals.size):
        for j in range(0, theta1_vals.size):
            t = np.array([[theta0_vals[i]], [theta1_vals[j]]])
            J_vals[i][j] = computeCost(X, Y, t)

    xs,ys = np.meshgrid(theta0_vals,theta1_vals)
    J_vals = np.transpose(J_vals)

    fig1 = plt.figure(1)
    ax = fig1.gca(projection="3d")
    ax.plot_surface(xs, ys, J_vals)
    plt.xlabel(r'$\theta_0$')
    plt.ylabel(r'$\theta_1$')

    plt.figure(2)
    lvls = np.logspace(-2, 3, 20)
    plt.contour(xs, ys, J_vals, levels=lvls, norm=LogNorm())
    plt.plot(theta[0], theta[1], c='r', marker="x")
    plt.show()