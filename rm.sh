#!/bin/bash

if [-f main];
then
	rm main
	echo "rm main"
fi

if [-f myTodoList];
then
	rm myTodoList;
fi
