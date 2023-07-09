import keras


class Prevue:
    def __init__(
        self, connector_name: str, project_name: str, uid: str, url: str
    ) -> None:
        """Define the initial variables for connection.

        Args:
            connector_name (str): Connector name.
            project_name (str): Project name.
            uid (str): User id. 
            url (str): Url. 
        """
        self.connector_name = connector_name
        self.project_name = project_name
        self.uid = uid
        self.url = url

        # NOTE add user name or authenticator

    def capture(self, metrics: dict):
        """Capture metrics.

        Args:
            metrics (dict): Metrics defined in the dict format.
        """

        print(metrics)


class PrevueKerasCallback(keras.callbacks.Callback, Prevue):
    def __init__(self, connector_name: str, project_name: str, uid: str, url: str):
        Prevue.__init__(self, connector_name, project_name, uid, url)

    def on_train_batch_end(self, batch, logs=None):
        self.capture({"batch": batch, "loss": logs["loss"]})

    def on_test_batch_end(self, batch, logs=None):
        self.capture({"batch": batch, "loss": logs["loss"]})

    def on_epoch_end(self, epoch, logs=None):
        self.capture(
            {"epoch": epoch, "loss": logs["loss"], "accuracy": logs["accuracy"]}
        )
