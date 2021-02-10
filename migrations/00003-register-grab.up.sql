insert into maintainer
(
    handle,
    name,
    email
)
values 
(
    'sharpvik',
    'Viktor A. Rozenko Voitenko',
    'sharp.vik@gmail.com'
)
on conflict do nothing;

insert into package
(
    name,
    maintainer,
    repository
)
values
(
    'grab-server',
    'sharpvik',
    'github.com/grab-pack/server'
)
on conflict do nothing;
