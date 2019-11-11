# coding=utf-8
import numpy as np


def featureNormalize(X):
    """
    Normalizes the features in X
    :param X:
    :return:
    """
    mu = np.mean(X, axis=0)
    sigma = np.std(X, axis=0, ddof=1)
    X_norm = (X - mu) / sigma
    return X_norm, mu, sigma


def computeCostMulti(X, y, theta):
    """
    Compute cost for linear regression with multiple variables
    :param X:
    :param y:
    :param theta:
    :return:
    """
    cost = np.dot(X, theta) - y
    return np.sum(np.power(cost, 2)) / (2 * y.shape[0])


def gradientDescentMulti(X, y, theta, alpha, num_iters):
    """
    Performs gradient descent to learn theta
    :param X:
    :param y:
    :param theta:
    :param alpha:
    :param num_iters:
    :return:
    """
    m = y.shape[0]
    J_history = np.zeros((num_iters, 1))
    O_theta = theta
    for i in range(num_iters):
        bias = np.dot(np.transpose(X), np.dot(X, O_theta) - y)
        O_theta = O_theta - (alpha / m) * bias
        J_history[i] = computeCostMulti(X, y, O_theta)
    return O_theta, J_history


def normalEqn(X, y):
    """
    Computes the closed-form solution to linear regression
    :param X:
    :param y:
    :return:
    """
    theta = np.linalg.pinv(X.T.dot(X)).dot(X.T).dot(y)
    return theta


if __name__ == "__main__":
    # ================ Part 1: Feature Normalization ================
    print('Loading data ...')
    data = np.loadtxt(r"ex1data2.txt", delimiter=",")
    X = data[:, 0:2]
    y = data[:, 2]
    m = y.shape[0]
    Y = y.reshape((m, 1))
    # Print out some data points
    print('First 10 examples from the dataset: ')
    print(' x = \n{},\n Y = \n{} '.format(X[0:10], Y[0:10]))

    # Scale features and set them to zero mean
    print('Normalizing Features ...')
    X, mu, sigma = featureNormalize(X)

    # Add intercept term to X
    X = np.c_[np.ones((m, 1)), X]

    # ================ Part 2: Gradient Descent ================
    print('Running gradient descent ...')

    # Choose some alpha value
    alpha = 0.1
    num_iters = 400

    # Init Theta and Run Gradient Descent
    theta = np.zeros((3, 1))
    theta, J_history = gradientDescentMulti(X, Y, theta, alpha, num_iters)
    # # Plot the convergence graph
    # plt.figure(0)
    # plt.plot(range(1, num_iters+1), J_history, color='b')
    # plt.xlabel("Number of iterations")
    # plt.ylabel("Cost J")
    # plt.show()

    # Display gradient descent's result
    print('Theta computed from gradient descent: ', theta)

    # Estimate the price of a 1650 sq-ft, 3 br house
    # Recall that the first column of X is all-ones. Thus, it does not need to be normalized.
    price = np.dot(np.array([1, (1650 - mu[0]) / sigma[0], (3 - mu[1]) / sigma[1]]), theta)
    print('Predicted price of a 1650 sq-ft, 3 br house (using gradient descent):', price)

    # ================ Part 3: Normal Equations ================
    print('Solving with normal equations...')
    data = np.loadtxt(r"ex1data2.txt", delimiter=",")
    X = data[:, 0:2]
    y = data[:, 2]
    m = y.shape[0]
    X = np.hstack((np.ones((m, 1)), X))
    Y = y.reshape((m, 1))
    theta = normalEqn(X, Y)
    # Display normal equation's result
    print('Theta computed from the normal equations: ', theta)
    # Estimate the price of a 1650 sq-ft, 3 br house
    price = np.array([1, 1650, 3]).dot(theta)
    print('Predicted price of a 1650 sq-ft, 3 br house:', price)
