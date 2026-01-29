global _start

section .text
_start:
    ; write syscall
    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    mov rsi, message    ; message to print
    mov rdx, 13         ; length of message
    syscall

    ; exit syscall
    mov rax, 60         ; sys_exit
    xor rdi, rdi        ; exit code 0
    syscall

section .data
message: db "Hello, World", 10  ; 10 = newline
