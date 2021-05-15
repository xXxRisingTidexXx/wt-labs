FROM migrate/migrate:v4.12.2
ADD https://raw.githubusercontent.com/eficode/wait-for/master/wait-for .
RUN chmod +x wait-for
