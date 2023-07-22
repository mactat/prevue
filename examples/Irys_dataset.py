import pandas as pd
import numpy as np
import tensorflow as tf
from sklearn.datasets import load_iris
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
import sys

sys.path.append("../")
from connectors.exporter import PrevueKerasCallback, Prevue


# #%%
data = load_iris()
# print the description of data
print(data.DESCR)
df = pd.DataFrame(data.data, columns = data.feature_names)
df['target'] = data.target
target_names = dict(zip((0,1,2), data.target_names))

#%%
Y = df['target']
X = df.drop('target', axis = 1)

#%%

seed = 42
np.random.seed(seed)

Y_final = tf.keras.utils.to_categorical(Y)
x_train, x_test, y_train, y_test = train_test_split(X, Y_final, test_size=0.25, random_state=seed, stratify=Y, shuffle=True)

print("Training Input shape\t: {}".format(x_train.shape))
print("Testing Input shape\t: {}".format(x_test.shape))
print("Training Output shape\t: {}".format(y_train.shape))
print("Testing Output shape\t: {}".format(y_test.shape))

#%%

std_clf = StandardScaler()
x_train_new = std_clf.fit_transform(x_train)
x_test_new = std_clf.transform(x_test)

#%%

model = tf.keras.models.Sequential()
model.add(tf.keras.layers.Dense(10, input_dim=4, activation=tf.nn.relu, kernel_initializer='he_normal',
                                kernel_regularizer=tf.keras.regularizers.l2(0.01)))
model.add(tf.keras.layers.BatchNormalization())
model.add(tf.keras.layers.Dropout(0.3))
model.add(tf.keras.layers.Dense(7, activation=tf.nn.relu, kernel_initializer='he_normal',
                                kernel_regularizer=tf.keras.regularizers.l1_l2(l1=0.001, l2=0.001)))
model.add(tf.keras.layers.BatchNormalization())
model.add(tf.keras.layers.Dropout(0.3))
model.add(tf.keras.layers.Dense(5, activation=tf.nn.relu, kernel_initializer='he_normal',
                                kernel_regularizer=tf.keras.regularizers.l1_l2(l1=0.001, l2=0.001)))
model.add(tf.keras.layers.Dense(3, activation=tf.nn.softmax))

model.compile(optimizer='adam', loss='categorical_crossentropy', metrics=['accuracy'])


prevueCallback = PrevueKerasCallback(
    user_id="maciek2",
    url="localhost:8080",
    email="maciek2@gmail.com",
    passwoard="MaciekToKozak",
    connector_name="keras",
    project_name="iris",
    model_name="modeltest",
)


iris_model = model.fit(x_train_new, y_train, epochs=10, batch_size=7, callbacks = [prevueCallback])

#%%

model.evaluate(x_test_new, y_test)
